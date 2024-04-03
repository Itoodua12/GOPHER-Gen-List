package main

import (
	"fmt"
	"reflect"

	hash_table "github.com/Itoodua12/GOPHER-Gen-List/internal/hash-table"
	linkedlist "github.com/Itoodua12/GOPHER-Gen-List/internal/linked-list"
	"github.com/Itoodua12/GOPHER-Gen-List/internal/queue"
	"github.com/Itoodua12/GOPHER-Gen-List/internal/set"
	"github.com/Itoodua12/GOPHER-Gen-List/internal/stack"
)

func main() {


	fmt.Println(" ##### HASH TABLE ##### ")
	test_hashtable()

	fmt.Println(" ##### STACK ##### ")
	stack := stack.Stack{}

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

	test_linked_list()


}

func test_hashtable() {
	hashTable := hash_table.Init()
	list := []string{
		"D",
		"E",
		"A",
		"B",
		"C",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	fmt.Println(hashTable.Search("A"))
	hashTable.Delete("A")
	fmt.Println(hashTable.Search("A"))
}


func test_linked_list() {
	fmt.Println(" ##### LINKED LIST ##### ")
	l := linkedlist.LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)

	get := l.Get(2)
	fmt.Println("there is : ", get)
	l.Delete(2)
	l.String()
}