package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Trie struct {
	Root *Node
}

type Node struct {
	Children map[byte]*Node
	Value    byte
}

func (trie Trie) Find(value string) bool {
	node := trie.Root

	for _, chr := range value {
		if newNode, ok := node.Children[byte(chr)]; ok {
			node = newNode
		} else {
			return false
		}
	}

	return true
}

func (trie Trie) Insert(value string) {
	node := trie.Root
	index := 0
	length := len(value)

	// Get the known nodes
	for index < length {
		if newNode, ok := node.Children[value[index]]; ok {
			node = newNode
			index += 1
		} else {
			break
		}
	}

	for index < length {
		val := value[index]
		node.Children[val] = &Node{
			Children: make(map[byte]*Node),
			Value:    val,
		}
		index += 1
	}
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			Children: make(map[byte]*Node),
			Value:    byte('r'),
		},
	}
}

func testTrie(trie *Trie, word string) {
	if trie.Find(word) {
		fmt.Println(word, "is a word")
	} else {
		fmt.Println(word, "is not a word")
	}
}

func main() {
	dictTrie := NewTrie()
	wordsFilePath := "/usr/share/dict/words"

	fmt.Println("Loading: ", wordsFilePath)

	contents, _ := ioutil.ReadFile(wordsFilePath)
	words := strings.Split(string(contents), "\n")

	for _, word := range words {
		dictTrie.Insert(strings.ToLower(word))
	}

	// Loltests

	testTrie(dictTrie, "zoo")
	testTrie(dictTrie, "zoom")
	testTrie(dictTrie, "alpha")
	testTrie(dictTrie, "megazoom")
}
