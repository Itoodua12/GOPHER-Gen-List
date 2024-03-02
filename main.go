package main

import (
	"fmt"
	"genericlist/genericlist"
)

func main() {

	list := genericlist.Create[string]()

	list.Add("El Matador")
	list.Add("The Machine")
	list.Add("Jon Bones Jones")
	list.RemoveByValue("Jon Bones Jones")
	list.Add("Volk")
	list.Remove(1)

	fmt.Println(list.Get(0)) // El Matador

	list.Print()

}
