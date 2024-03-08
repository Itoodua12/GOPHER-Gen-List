package queue

import "errors"


// It holds elements of any type. 

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, errors.New("Queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, errors.New("Queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
