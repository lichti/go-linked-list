package linkedlist

import (
	"bufio"
	"fmt"
	"os"
)

// Item .
type Item struct {
	Value interface{}
	Next  *Item
}

// List .
type List struct {
	First *Item
}

// Len return size of list
func (l *List) Len() int {
	p := l.First
	var i int
	i = 0

	for p.Next != nil {
		p = p.Next
		i++
		if p.Next == nil {
			i++
		}
	}
	return i
}

// Push a new item into list
func (l *List) Push(value interface{}) {
	n := &Item{value, nil}
	if l.First == nil {
		l.First = n
	} else {
		where := l.First
		for where.Next != nil {
			where = where.Next
		}
		where.Next = n
	}
}

// Find a Item into List
func (l *List) Find(value interface{}) bool {
	var p *Item
	if l.First == nil {
		return false
	}
	p = l.First
	for p.Next != nil {
		if p.Value == value {
			return true
		}
		p = p.Next
	}
	return false
}

// Remove item of list and return a bool value
func (l *List) Remove(value interface{}) bool {
	var p *Item
	if l.Find(value) == false {
		return false
	}
	for l.First.Value == value {
		l.First = l.First.Next
	}
	if l.First == nil {
		return false
	}
	p = l.First
	for p != nil {
		if p.Next != nil {
			if p.Next.Value == value {
				p.Next = p.Next.Next
			}
		}
		p = p.Next
	}
	return true
}

// Print List to stdOut
func (l *List) Print() {
	p := l.First
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}
}

// Load list from file
func (l *List) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.Push(scanner.Text())
	}
	return scanner.Err()
}

// Save linked list to file
func (l *List) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	p := l.First
	for i := 0; i < l.Len(); i++ {
		fmt.Fprintln(w, p.Value)
		p = p.Next
	}
	return w.Flush()
}

// Invert linked list
func (l *List) Inversor() {
	var prev *Item
	curr := l.First
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.First = prev
}

// Copy itens from other linked list
func (l *List) Copy(origin List) {
	p := origin.First
	for p.Next != nil {
		l.Push(p.Value)
		p = p.Next
		if p.Next == nil {
			l.Push(p.Value)
		}
	}
}
