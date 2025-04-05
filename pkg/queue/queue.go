package queue

import "errors"

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
