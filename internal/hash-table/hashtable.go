package hash_table

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}

type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func (b *Bucket) insert(k string) {
	if !b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

func (b *Bucket) search(k string) bool {
	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

func (b *Bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prev := b.head
	for prev.next != nil {
		if prev.next.key == k {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}
