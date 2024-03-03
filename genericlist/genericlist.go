package genericlist

import "fmt"

type GenericList[T comparable] struct {
	data []T
}

func Create[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Add(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(index int) T {
	if index > len(l.data)-1 || index < 0 {
		panic("Array Index Out Of Bounds")
	}
	return l.data[index]
}

func (l *GenericList[T]) Remove(index int) {
	if index > len(l.data)-1 || index < 0 {
		panic("Array Index Out Of Bounds")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
}

func (l *GenericList[T]) RemoveByValue(value T) {
	for i, v := range l.data {
		if v == value {
			l.data = append(l.data[:i], l.data[i+1:]...)
			return
		}
	}
	panic("Input value does not exist in the List")
}

func (l *GenericList[T]) Print() {
	for _, v := range l.data {
		fmt.Println(v)
	}
}

func (l *GenericList[T]) Clear() {
	l.data = nil
}

func (l *GenericList[T]) IsEmpty() bool {
    return len(l.data) == 0
}