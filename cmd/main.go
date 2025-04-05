package main

import (
	"fmt"

	"github.com/SergeyGolang/stack/pkg/queue"
)

func main() {
	myQueue := queue.NewQueue()
	myQueue.Enqueue(5)
	myQueue.Enqueue("Hello")
	myQueue.Enqueue(10)

	item, err := myQueue.Front()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item)
	}

	item, err = myQueue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item)
	}

	item, err = myQueue.Front()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item)
	}

	fmt.Println(*myQueue)

}
