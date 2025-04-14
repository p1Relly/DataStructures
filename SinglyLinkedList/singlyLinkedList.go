package main

import (
	"fmt"
)

// Node - узел односвязного списка
type Node struct {
	Value int   // Значение
	Next  *Node // Указатель на следующий
}

// NewNode - конструктор узла
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// SinglyLinkedList - односвязный список
type SinglyLinkedList struct {
	Head *Node // Указатель на начало списка

	count int // Кол-во узлов в списке
}

// NewSinglyLinkedList - конструктор списка
func NewSinglyLinkedList(head *Node) *SinglyLinkedList {
	return &SinglyLinkedList{Head: head}
}

// GetCount - возвращает кол-во узлов списка
func (l *SinglyLinkedList) GetCount() int {
	return l.count
}

// Print - возвращает строковое представление списка
func (l *SinglyLinkedList) Print() string {
	var result string
	current := l.Head
	for current != nil {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next
	}

	return result
}

// Find - ищет узел списка по значению
func (l *SinglyLinkedList) Find(key int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}
	current := l.Head
	for current.Value != key {
		current = current.Next
		if current == nil {
			return nil
		}
	}
	return current
}

// FindByIndex - ищет узел списка по индексу
func (l *SinglyLinkedList) FindByIndex(index int) *Node {
	if index >= l.GetCount() || index < 0 || l.Head == nil {
		return nil
	}

	current := l.Head

	for range index {
		current = current.Next
	}
	return current
}

// FindLast - ищет последний узел равный переданному значению
func (l *SinglyLinkedList) FindLast(key int) *Node {
	if l == nil || l.count == 0 {
		return nil
	}

	current := l.Head
	var lastFound *Node = nil

	for current != nil {
		if current.Value == key {
			lastFound = current
		}
		current = current.Next
	}
	return lastFound
}

// PushBack - добавляет узел в конец
func (l *SinglyLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if l.count == 0 {
		l.Head = newNode
	} else {
		tail := l.FindTail()
		tail.Next = newNode
	}
	l.count++
}

// FindTail - ищет последний узел списка
func (l *SinglyLinkedList) FindTail() *Node {
	if l.count == 0 {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// PushFront - добавляет узел в начало
func (l *SinglyLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if l.count != 0 {
		newNode.Next = l.Head
	}
	l.Head = newNode
	l.count++
}

// AddAfter - добавляет узел после заданного
func (l *SinglyLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	newNode := &Node{Value: item}
	newNode.Next = node.Next
	node.Next = newNode
	l.count++
}

// AddBefore - добавляет узел перед заданным
func (l *SinglyLinkedList) AddBefore(node *Node, item int) {
	if node == nil || l.Head == nil {
		return
	}
	if node == l.Head {
		newNode := &Node{Value: item, Next: l.Head}
		l.Head = newNode
		l.count++
		return
	} else {
		previous := l.Head
		for previous.Next != node {
			previous = previous.Next
		}
		newNode := &Node{Value: item, Next: node}
		previous.Next = newNode
		l.count++
	}
}

// PushBackRange - добавляет список узлов в конец
func (l *SinglyLinkedList) PushBackRange(items []int) {
	if len(items) == 0 {
		return
	}

	for _, v := range items {
		l.PushBack(v)
	}
}

// Удаление первого элемента.
func (l *SinglyLinkedList) RemoveFirst() {
	if l.count == 0 {
		return
	}
	l.Head = l.Head.Next
	l.count--
}

// Удаление последнего элемента
func (l *SinglyLinkedList) RemoveLast() {
	if l.count == 0 {
		return
	}
	if l.count == 1 {
		l.Head = nil
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	l.count--
}

// Удаление произвольного узла
func (l *SinglyLinkedList) RemoveNode(node *Node) {
	if l.Head == node {
		l.Head = node.Next
	} else {
		current := l.Head
		for current.Next != nil {
			if current.Next == node {
				break
			}
			current = current.Next
		}
		current.Next = node.Next
	}
	l.count--
}

func (l *SinglyLinkedList) Remove(key int) bool {
	if current := l.Find(key); current != nil {
		l.RemoveNode(current)
		return true
	}

	return false
}

// RemoveLastKey - удаляет последний узел по ключу
func (l *SinglyLinkedList) RemoveLastKey(key int) bool {
	if current := l.FindLast(key); current != nil {
		l.RemoveNode(current)
		return true
	}

	return false
}

// RemoveAll - удаляет все узлы по ключу
func (l *SinglyLinkedList) RemoveAll(key int) (count int) {
	for current := l.Head; current != nil; {
		if current.Value == key {
			l.RemoveNode(current)
			count++
		}
		current = current.Next
	}
	return
}

// Reverse - переворачивает все элементы в обратном порядке
func (l *SinglyLinkedList) Reverse() {
	list := make([]int, 0, l.count)
	current := l.Head

	for current != nil {
		list = append(list, current.Value)
		current = current.Next
	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	current = l.Head
	for _, value := range list {
		current.Value = value
		current = current.Next
	}
}

func main() {
	singlyLinkedList := NewSinglyLinkedList(nil)

	count := singlyLinkedList.GetCount() // Получить количество элементов
	fmt.Println(count)                   // Выводит 0

	node := singlyLinkedList.Find(7) // Получить узел, значение которого равно 7
	fmt.Println(node == nil)         // Выводит true
	node = singlyLinkedList.Find(1)  // Получить узел, значение которого равно 1
	fmt.Println(node == nil)         // Выводит true

	printedList := singlyLinkedList.Print() // Распечатать список
	fmt.Println(printedList)                // Выводит пустую строку
}
