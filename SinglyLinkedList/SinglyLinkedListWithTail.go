package main

import "fmt"

// Node - узел списка
type Node struct {
	Next  *Node
	Value int
}

// NewNode - конструктор узла
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// SinglyLinkedListWithTail - односвязный список с указателем на хвост списка
type SinglyLinkedListWithTail struct {
	Head  *Node
	Tail  *Node
	count int
}

// GetCount - возвращает кол-во узлов списка
func (list *SinglyLinkedListWithTail) GetCount() int {
	return list.count
}

// Print - возвращает строковое представление списка
func (list *SinglyLinkedListWithTail) Print() string {
	result := ""
	current := list.Head
	for current != nil {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next
	}
	return result
}

// Find - ищет узел по значению
func (list *SinglyLinkedListWithTail) Find(key int) *Node {
	if list.count == 0 {
		return nil
	}
	current := list.Head
	for current.Value != key {
		current = current.Next
		if current == nil {
			return nil
		}
	}
	return current
}

// PushBack - добавляет узел в конец
func (list *SinglyLinkedListWithTail) PushBack(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.Head = newNode
	} else {
		list.Tail.Next = newNode
	}
	list.count++
	list.Tail = newNode
}

// PushFront - добавляет узел в начало
func (list *SinglyLinkedListWithTail) PushFront(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
	}
	list.Head = newNode
	list.count++
}

// AddAfter - добавляет узел после заданного
func (list *SinglyLinkedListWithTail) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	newNode := &Node{Value: item}
	newNode.Next = node.Next
	node.Next = newNode
	if node == list.Tail {
		list.Tail = newNode
	}
	list.count++
}

// RemoveFirst - удаляет голову списка
func (list *SinglyLinkedListWithTail) RemoveFirst() {
	if list.count == 0 {
		return
	}
	if list.count == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		list.Head = list.Head.Next
	}
	list.count--
}

// RemoveLast - удаляет последний узел
func (list *SinglyLinkedListWithTail) RemoveLast() {
	if list.count == 0 {
		return
	}
	if list.count == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		current := list.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
		list.Tail = current
	}
	list.count--
}

// RemoveNode - удаляет произвольный узел
func (list *SinglyLinkedListWithTail) RemoveNode(node *Node) {
	if node == list.Head {
		list.RemoveFirst()
	} else {
		current := list.Head
		for current.Next != nil {
			if current.Next == node {
				break
			}
			current = current.Next
		}
		current.Next = node.Next
		if node == list.Tail {
			list.Tail = current
		}
		list.count--
	}
}
