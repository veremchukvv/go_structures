package main

import (
	"log"
)

type List struct {
	root *node
	head *node
	len  int
}

type node struct {
	data interface{}
	next *node
	prev *node
}

func (l *List) Push(data interface{}) {
	n := &node{
		data: data,
	}

	if l.root == nil && l.head == nil {
		l.root = n
		l.head = n
	} else {
		n.prev = l.head
		l.head.next = n
		l.head = n
	}
	l.len++
}

func (l *List) Pop() {
	if l.head.prev != nil {
		log.Println("popping!")
		l.head = l.head.prev
		l.head.next = nil
	}
}

func (l *List) Peek() {
	log.Println("peeking!", l.head.data)
}

func (l *List) ToArray() []interface{} {
	arr := make([]interface{}, 0, l.len)

	current := l.root
	for current != nil {
		arr = append(arr, current.data)
		current = current.next
	}

	return arr
}
