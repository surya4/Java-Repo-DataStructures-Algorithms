package main

import "fmt"

const SIZE = 15

type Node struct {
	val  int
	next *Node
}

type HashTable struct {
	table map[int] *Node
	size int
}

func hashFunction(i , size int) int {
	return i % size
}

func insert(hash *HashTable, data int) int {
	index := hashFunction(data, hash.size)
	element := &Node{val: data, next: hash.table[index]}
	hash.table[index] = element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.table {
		if hash.table[k] != nil {
			t := hash.table[k]
			for t != nil {
				fmt.Print(t.val, " ")
				t = t.next
			}
			fmt.Println()
		}
	}
}

func lookUp(hash *HashTable, data int) bool {
	index := hashFunction(data, hash.size)
	if hash.table[index] != nil {
		t := hash.table[index] 
		for t != nil {
			if t.val == data {
				return true
			}
			t = t.next
		}
	}
	return false
}


func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{table: table, size: SIZE}

	for i:=0 ; i< 120; i++ {
		insert(hash, i)
	}

	traverse(hash)
	fmt.Println("Is 78 availilable in hash table -> ", lookUp(hash, 78))
	fmt.Println("Is 178 availilable in hash table ->", lookUp(hash, 178))
}
