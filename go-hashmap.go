package main

import (
	"fmt"
	"hash/crc32"
)

const (
	HashSize = 10
)

func hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v < 0 {
		return -v
	}
	return v
}

type Entry struct {
	Key string
	Value interface{}
	Next *Entry
}

type HashMap struct {
	Size int
	Entries []*Entry
}

func (this *HashMap) Set(key string, value interface{}) {
	index := hashcode(key) % this.Size
	node := this.Entries[index]
	entry := &Entry{
		Key: key,
		Value: value,
	}

	if node == nil {
		this.Entries[index] = entry
	} else if node.Key == key { // In the case we want to overwrite the key's value
		node.Value = value
	}	else {
		for node.Next != nil {
			if node.Key == key { // In the case we want to overwrite the key's value
				node.Value = value
				return
			}
			node = node.Next
		}
		node.Next = entry
	}
}

func (this *HashMap) Get(key string) interface{} {
	index := hashcode(key) % this.Size
	node := this.Entries[index]

	if node == nil {
		return nil
	} else {
		for node != nil && node.Key != key {
			node = node.Next
		}
	}

	if node == nil {
		return nil
	} else {
		return node.Value
	}
}

func main() {

	root := &HashMap{
		Size: HashSize,
		Entries: make([]*Entry, HashSize, HashSize),
	}

	root.Set("hello", "world")
	root.Set("potato", "rocket")
	root.Set("potato", "potatoes")
	root.Set("world", 33)
	root.Set("woo", 44.4)

	fmt.Println(root.Get("potato"))
	fmt.Println(root.Get("lol"))
	fmt.Println(root.Get("woo"))
}
