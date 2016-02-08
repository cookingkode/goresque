package goresque

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
)

type job struct {
	Class string        `json:"class"`
	Args  []interface{} `json:"args"`
}

func DoInit(redisAddress, redisPassword string) {
	pool = newPool(redisAddress, redisPassword)

}

func AddJob(queue, jobClass string, args ...interface{}) (int64, error) {
	conn := pool.Get()
	defer conn.Close()

	// NOTE: Dirty hack to make a [{}] JSON struct
	if len(args) == 0 {
		args = append(make([]interface{}, 0), make(map[string]interface{}, 0))
	}

	jobJSON, err := json.Marshal(&job{jobClass, args})
	if err != nil {
		return -1, err
	}

	resp, err := conn.Do("RPUSH", queue, string(jobJSON))

	return redis.Int64(resp, err)

}

var (
	pool *redis.Pool
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
