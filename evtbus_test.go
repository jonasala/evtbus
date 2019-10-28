package evtbus_test

import (
	"fmt"

	"github.com/jonasala/evtbus"
)

//EventBus ideally must be in global scope
var eb = evtbus.New()

func Example() {

	//Creating channels for receive events
	ch1 := make(evtbus.EventChannel)
	ch2 := make(evtbus.EventChannel)
	ch3 := make(evtbus.EventChannel)

	//Subscribe channels to topics
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic1", ch2)
	eb.Subscribe("topic2", ch3)

	//Publish an event to a topic. Data is interface{} so it can be anything
	eb.Publish(evtbus.Event{
		Topic: "topic1",
		Data:  "Published on topic1",
	})

	//Publish an event to a topic. Data is interface{} so it can be anything
	eb.Publish(evtbus.Event{
		Topic: "topic2",
		Data:  "Published on topic2",
	})

	//Receive events on channels
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)
	// Output:
	// {topic1 Published on topic1}
	// {topic1 Published on topic1}
	// {topic2 Published on topic2}
}
