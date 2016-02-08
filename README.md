# goresque
Golang Resque Client 

## Need
Built primarily to work  as a client for  [Goworker](https://www.goworker.org/)


## Usage
Heres a sample usecase of initiatlizing a client and queing two jobs
The client is thread-sage

``` go

	import "github.com/cookingkode/goresque"
	
	brokerAddr := ":6379" //local redis instance
	namespace := "name_of_my_app"
	queue := "high_priority_things"

 	client := goresque.DoInit(brokerAddr,  namespace, queue)

 	client.AddJob("SomeJobClass", "InputA", "InputB")
	client.AddJob("SomeOtherJobClass", "InputA", "InputB", "InputB")
```

This assumes that you have registed some job classes  with  goworker.Register. Example


``` go
	goworker.Register("SomeJobClass", DoSomething)

	func DoSomething(queue string, args ...interface{}) error {
		InputA := args[0]
		InputB := args[1]

		// do something useful

	}

```
