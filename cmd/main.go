package main

import (
	"fmt"
)

/*
	Hash table collision is when two keys have the same hash code.
	Two ways to handle hash table collision is:
	1) open addressing (putting the key in the next available slot in array)
	2) seperate chaining

	for open addressing as more and more items are adding and get further and further away from the expected index we lose benefits
	seperate chaining is storing multiple names in one slot and this is done by making the type of the index a linked list.

	Methods of HashTables
	1) Search
	2) Insert
	3) Delete

	Time Complexity of Hash Tables is constant on an average case worst case would be O(n)
	1) Search O(1)
	2) Insert O(1)
	3) Delete O(1)

*/

// ArraySize is the size of the hashtable array
const ArraySize = 7

// HashTable structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure is a linked list
type bucket struct {
	head *bucketNode
}

// bucketNode structure is a linked list node that holds
type bucketNode struct {
	key  string
	next *bucketNode
}

// HashTable methods
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// // Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// // Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
	return
}

// Bucket methods
// Insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("The key: %s already exist", k)
	}
}

// Search
func (b *bucket) search(k string) bool {
	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	current := b.head
	for current.next != nil {
		if current.next.key == k {
			current.next = current.next.next
		}
		current = current.next
	}
	return
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	if hashTable.Search("STAN") {
		fmt.Println("STAN exist failed to remove stan.")
	} else {
		fmt.Println("STAN does not exist successfully removed.")
	}
}
