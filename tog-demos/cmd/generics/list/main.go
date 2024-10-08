package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Prepend(elem T) {
	tail := *l
	l.next = &tail
	l.val = elem
}

func (l *List[T]) LastElem() *List[T] {
	for l.next != nil {
		l = l.next
	}
	return l
}

func (l *List[T]) Add(elem T) {
	last := l.LastElem()
	last.next = &List[T]{nil, elem}
}

func (l *List[T]) String() string {
	result := ""
	for l != nil {
		val := fmt.Sprintf("| %v ", l.val)
		result += val
		l = l.next
	}
	return string(result)
}

func main() {
	l := List[int]{nil, 42}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Prepend(12)
	fmt.Printf("%#v\n", l)
	fmt.Printf("%s\n", l.String())
}
