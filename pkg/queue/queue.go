package queue

import "errors"

/*
A queue implements FIFO (Fast In, First Out).
A queue has four methods:
 1. Enqueue - append a value
 2. Dequeue - return and remove the last element
 3. Front - retrieve the last element
 4. IsEmpty - check stack if the stack is empty
*/

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	item := q.items[0]

	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
