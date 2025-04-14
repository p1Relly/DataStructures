package main

import (
	"fmt"
	"strings"
)

// Node - узел списка
type Node struct {
	Next     *Node
	Previous *Node
	Value    int
}

// NewNode - конструктор узла
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// DoubleCircularLinkedList - двусвязный циклический список
type DoubleCircularLinkedList struct {
	Head  *Node
	count int
}

// GetCount - возвращает размер списка
func (list *DoubleCircularLinkedList) GetCount() int {
	return list.count
}

// Print - возвращает строковое представление списка
func (list *DoubleCircularLinkedList) Print(forward bool) string {
	if list.count == 0 {
		return ""
	}
	if forward {
		return list.print(list.Head, forward)
	}
	return list.print(list.Head.Previous, forward)
}

// print - возвращает строковое представление списка начиная с переданного узла
func (list *DoubleCircularLinkedList) print(node *Node, forward bool) string {
	if list.count == 0 {
		return ""
	}
	var result strings.Builder
	current := node
	for {
		result.WriteString(fmt.Sprintf("%d ", current.Value))
		if forward {
			current = current.Next
		} else {
			current = current.Previous
		}
		if current == node {
			break
		}
	}
	return result.String()
}

// Find - поиск узла по значению
func (list *DoubleCircularLinkedList) Find(key int) *Node {
	if list.count == 0 {
		return nil
	}
	current := list.Head
	for {
		if current.Value == key {
			return current
		}
		current = current.Next
		if current == list.Head {
			break
		}
	}
	return nil
}

// FindLast - находит последний узел по значению
func (l *DoubleCircularLinkedList) FindLast(key int) *Node {
	if l.count == 0 {
		return nil
	}

	current := l.Head.Previous
	for range l.count {
		if current.Value == key {
			return current
		}
		current = current.Previous
	}
	return nil
}

func (list *DoubleCircularLinkedList) addBeforeInternal(node, newNode *Node) {
	newNode.Next = node
	newNode.Previous = node.Previous
	node.Previous.Next = newNode
	node.Previous = newNode
	list.count++
}

func (l *DoubleCircularLinkedList) insertNodeToEmptyList(node *Node) {
	node.Next = node
	node.Previous = node
	l.Head = node
	l.count++
}

// AddBefore - вставляет новый узел перед переданным узлом
func (list *DoubleCircularLinkedList) AddBefore(node *Node, item int) {
	newNode := &Node{Value: item}
	list.addBeforeInternal(node, newNode)
	if node == list.Head {
		list.Head = newNode
	}
}

// PushBack - добавляет узел в конец
func (list *DoubleCircularLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addBeforeInternal(list.Head, newNode)
	}
}

// PushFront - добавление узла в начало списка
func (list *DoubleCircularLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addBeforeInternal(list.Head, newNode)
		list.Head = newNode
	}
}

func main() {
	doubleCircularLinkedList := DoubleCircularLinkedList{}
	doubleCircularLinkedList.PushBack(1)
	doubleCircularLinkedList.PushBack(2)
	doubleCircularLinkedList.PushBack(5)
	doubleCircularLinkedList.PushBack(2)
	doubleCircularLinkedList.PushBack(1)

	fmt.Println(doubleCircularLinkedList.print(doubleCircularLinkedList.Head, true))
	test := doubleCircularLinkedList.FindLast(3)
	fmt.Println(test)
}
