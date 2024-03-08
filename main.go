package main

import (
	"fmt"
	"genericlist/genericstack"
	"reflect"
)

func main() {

	stack := genericstack.Stack{}

	stack.Push(12)
	stack.Push("Any type u want")
	stack.Push(true)
	stack.Push(3.14159)
	stack.Push(func() {
		fmt.Println("This is Function !!!!")
	})

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		if reflect.TypeOf(item).Kind() == reflect.Func {
			fn, _ := item.(func())
			fn()
		} else {
			fmt.Printf("Value is -> %v\n", item)
		}
	}

}
