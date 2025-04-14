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

// SinglyCircularLinkedList - односвязный циклический список
type SinglyCircularLinkedList struct {
	Tail  *Node
	count int
}

// GetCount - возвращает кол-во узлов списка
func (list *SinglyCircularLinkedList) GetCount() int {
	return list.count
}

// Print - выводит строковое представление списка
func (list *SinglyCircularLinkedList) Print() string {
	if list.count == 0 {
		return ""
	}
	return list.PrintNode(list.Tail.Next)
}

// PrintNode - выводит строковое представление списка начиная с переданного узла
func (list *SinglyCircularLinkedList) PrintNode(node *Node) string {
	if list.count == 0 {
		return ""
	}
	result := ""
	current := node
	for {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next

		if current == node {
			break
		}
	}
	return result
}

// Добавление элемента в конец
func (list *SinglyCircularLinkedList) addAfterInternal(node, newNode *Node) {
	newNode.Next = node.Next
	node.Next = newNode
	list.count++
}

// Вставляет новый узел в пустой список
func (list *SinglyCircularLinkedList) insertNodeToEmptyList(node *Node) {
	node.Next = node
	list.Tail = node
	list.count++
}

// PushBack - добавляет узел в конец списка
func (list *SinglyCircularLinkedList) PushBack(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addAfterInternal(list.Tail, newNode)
		list.Tail = newNode
	}
}

// AddAfter - добавляет узел после заданного
func (list *SinglyCircularLinkedList) AddAfter(node *Node, item int) {
	newNode := &Node{Value: item}
	list.addAfterInternal(node, newNode)
	if node == list.Tail {
		list.Tail = newNode
	}
}

// PushFront - добавляет узел в начало списка
func (list *SinglyCircularLinkedList) PushFront(item int) {
	newNode := &Node{Value: item}
	if list.count == 0 {
		list.insertNodeToEmptyList(newNode)
	} else {
		list.addAfterInternal(list.Tail, newNode)
	}
}

// AddBefore - добавляет узел до переданного
func (l *SinglyCircularLinkedList) AddBefore(node *Node, item int) {
	if l.count == 0 {
		return
	}
	newNode := NewNode(item)
	current := l.Tail
	for current.Next != l.Tail {
		if current.Next == node {
			break
		}
		current = current.Next

	}
	newNode.Next = current.Next
	current.Next = newNode
	l.count++
}

// Find - ищет узел по значению
func (list *SinglyCircularLinkedList) Find(key int) *Node {
	if list.count == 0 {
		return nil
	}

	current := list.Tail
	for {
		if current.Value == key {
			return current
		}
		current = current.Next

		if current == list.Tail {
			break
		}
	}
	return nil
}

// FindLast - ищет последний узел в списке
func (list *SinglyCircularLinkedList) FindLast(item int) *Node {
	if list.count == 0 {
		return nil
	}

	var lastFound *Node
	current := list.Tail.Next // Начало списка (Head)
	start := current

	for {
		if current.Value == item {
			lastFound = current
		}

		// Переход к следующему узлу
		current = current.Next

		// Проверяем завершение полного обхода
		if current == start {
			break
		}
	}

	return lastFound
}

// Удаление элемента
func (list *SinglyCircularLinkedList) removeAfterNodeInternal(node *Node) {
	if list.count == 0 {
		return
	}
	if list.count == 1 {
		list.Tail = nil
	} else {
		if node.Next == list.Tail {
			list.Tail = node
		}
		node.Next = node.Next.Next
	}
	list.count--
}

// RemoveFirst - удаляет первый узел в списке
func (list *SinglyCircularLinkedList) RemoveFirst() {
	list.removeAfterNodeInternal(list.Tail)
}

// RemoveNode - удаляет переданный узел из списка
func (list *SinglyCircularLinkedList) RemoveNode(node *Node) {
	if list.count == 0 {
		return
	}
	current := node
	for current.Next != node {
		current = current.Next
	}
	list.removeAfterNodeInternal(current)
}

// RemoveLast - удаляет последний узел списка
func (list *SinglyCircularLinkedList) RemoveLast() {
	list.RemoveNode(list.Tail)
}

// Remove - удаляет первое вхождение указанного элемента из списка
func (l *SinglyCircularLinkedList) Remove(key int) bool {
	if node := l.Find(key); node != nil {
		l.RemoveNode(node)
		return true
	}

	return false
}

// RemoveLastKey - удаляет последний узел по ключу
func (l *SinglyCircularLinkedList) RemoveLastKey(key int) bool {
	if node := l.FindLast(key); node != nil {
		l.RemoveNode(node)
		return true
	}

	return false
}

func main() {
	singlyCircularLinkedList := &SinglyCircularLinkedList{}

	singlyCircularLinkedList.PushBack(1)
	singlyCircularLinkedList.PushBack(2)
	singlyCircularLinkedList.PushBack(3)
	singlyCircularLinkedList.PushBack(4)
	singlyCircularLinkedList.PushBack(5)
	fmt.Println(singlyCircularLinkedList.Print())
	singlyCircularLinkedList.AddBefore(&Node{Value: 3}, 6)
	fmt.Println(singlyCircularLinkedList.Print())
}
