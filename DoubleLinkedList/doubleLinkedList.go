package main

import (
	"errors"
	"fmt"
)

// Node - узел списка
type Node struct {
	Next     *Node // следующий узел
	Previous *Node // предыдущий узел
	Value    int   // значение
}

// NewNode - конструктор узла
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// DoubleLinkedList - двусвязный список
type DoubleLinkedList struct {
	Head  *Node
	count int
}

// GetCount - возвращает кол-во узлов списка
func (dll *DoubleLinkedList) GetCount() int {
	return dll.count
}

// Print - строковое представление списка
func (dll *DoubleLinkedList) Print() string {
	result := ""
	current := dll.Head
	for current != nil {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next
	}
	return result
}

// Find - ищет узел по значению
func (dll *DoubleLinkedList) Find(key int) *Node {
	if dll.count == 0 {
		return nil
	}

	current := dll.Head
	for current.Value != key {
		current = current.Next
		if current == nil {
			return nil
		}
	}
	return current
}

// FindLast - ищет последний узел по значению
func (dll *DoubleLinkedList) FindLast(key int) *Node {
	if dll.count == 0 || dll == nil {
		return nil
	}

	current := dll.Head
	var lastFound *Node
	for range dll.count {
		if current.Value == key {
			lastFound = current
		}
		current = current.Next
	}
	return lastFound
}

// PushBack - добавляет узел в конец списка
func (dll *DoubleLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if dll.count == 0 {
		dll.Head = newNode
	} else {
		tail := dll.findTail()
		tail.Next = newNode
		newNode.Previous = tail
	}
	dll.count++
}

// findTail - возвращает последний узел списка
func (dll *DoubleLinkedList) findTail() *Node {
	if dll.count == 0 {
		return nil
	}
	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// PushFront - добавляет узел в начало списка
func (dll *DoubleLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if dll.count != 0 {
		newNode.Next = dll.Head
		dll.Head.Previous = newNode
	}
	dll.Head = newNode
	dll.count++
}

// AddBefore - вставляет узел в список перед переданным
func (dll *DoubleLinkedList) AddBefore(node *Node, item int) {
	if node == nil {
		return
	}

	if node == dll.Head {
		dll.PushFront(item)
		return
	}
	newNode := &Node{
		Value:    item,
		Next:     node,
		Previous: node.Previous,
	}
	newNode.Previous.Next = newNode
	node.Previous = newNode
	dll.count++
}

// AddAfter - вставляет узел в список после переданного
func (dll *DoubleLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}

	newNode := &Node{
		Value:    item,
		Previous: node,
		Next:     node.Next,
	}

	if node.Next != nil {
		node.Next.Previous = newNode
	}

	node.Next = newNode

	dll.count++
}

// RemoveFirst - удаляет первый (голову) элемент списка
func (dll *DoubleLinkedList) RemoveFirst() error {
	if dll.count == 0 {
		return errors.New("Список пуст")
	}
	dll.Head = dll.Head.Next
	if dll.Head != nil {
		dll.Head.Previous = nil
	}
	dll.count--
	return nil
}

// RemoveLast - удаляет последний элемент списка
func (dll *DoubleLinkedList) RemoveLast() error {
	if dll.count == 0 {
		return errors.New("Список пуст")
	}
	if dll.count == 1 {
		dll.Head = nil
	} else {
		current := dll.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	dll.count--
	return nil
}

// RemoveNode - удаляет конкретный узел списка
func (dll *DoubleLinkedList) RemoveNode(node *Node) error {
	if dll.Head == node {
		return dll.RemoveFirst()
	}
	prev := node.Previous
	prev.Next = node.Next
	if node.Next != nil {
		node.Next.Previous = prev
	}
	dll.count--
	return nil
}

// RemoveFirstByValue - удаляет первый узел из списка, равный конкретному значению
func (dll *DoubleLinkedList) RemoveFirstByValue(value int) bool {
	if node := dll.Find(value); node != nil {
		dll.RemoveNode(node)
		return true
	}

	return false
}

// RemoveLastByValue - удаляет последний узел из списка, равный конкретному значению
func (dll *DoubleLinkedList) RemoveLastByValue(value int) bool {
	if node := dll.FindLast(value); node != nil {
		dll.RemoveNode(node)
		return true
	}

	return false
}

// Group - группирует узлы списка так, что сначала идут узлы не четных позиций, затем четных
func (dll *DoubleLinkedList) Group() {
	if dll.count < 2 {
		return
	}

	originalCount := dll.count
	current := dll.Head.Next

	for i := 1; i < originalCount; i += 2 {
		nextNode := current.Next

		dll.PushBack(current.Value)
		dll.RemoveNode(current)

		if nextNode != nil && nextNode.Next != nil {
			current = nextNode.Next
		} else {
			break
		}
	}
}

func main() {
	doubleLinkedList := &DoubleLinkedList{}

	doubleLinkedList.PushBack(8)
	doubleLinkedList.PushBack(1)
	doubleLinkedList.PushBack(6)
	doubleLinkedList.PushBack(4)
	doubleLinkedList.PushBack(2)
	doubleLinkedList.PushBack(2)
	doubleLinkedList.PushBack(3)
	doubleLinkedList.PushBack(7)
	doubleLinkedList.PushBack(5)

	printedList := doubleLinkedList.Print()
	fmt.Println(printedList)

	doubleLinkedList.Group()

	printedList = doubleLinkedList.Print()
	fmt.Println(printedList)

}
