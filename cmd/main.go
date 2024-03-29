package main

import (
	"fmt"
	"reflect"

	"github.com/Itoodua12/GOPHER-Gen-List/internal/genericstack"
	"github.com/Itoodua12/GOPHER-Gen-List/internal/queue"
	"github.com/Itoodua12/GOPHER-Gen-List/internal/set"
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
		el, _ := stack.Pop()
		if reflect.TypeOf(el).Kind() == reflect.Func {
			fn, _ := el.(func())
			fn()
		} else {
			fmt.Printf("Value is -> %v\n", el)
		}
	}

	fmt.Println(" ##### QUEUE ##### ")

	queue := queue.Queue{}

	queue.Enqueue(13)
	queue.Enqueue("GOpher")
	queue.Enqueue(false)

	for !queue.IsEmpty() {
		el, _ := queue.Dequeue()
		fmt.Printf("Value is -> %v\n", el)
	}

	fmt.Println(" ##### SET ##### ")

	set := set.NewMySet()

	set.Add("IT12")
	set.Add("IT12")
	set.Add(12)
	set.Add(12)

	fmt.Printf("set size is %v\n", set.Size()) // 2

	exists := set.Contains("IT12")
	fmt.Println("", exists)

}
