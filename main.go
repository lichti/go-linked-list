// Package linkedlist must help you  to create a linked list and provide some methods to easy manipule this list.

package linkedlist

import (
	"bufio"
	"fmt"
	"os"
)

// List is the main object this library, it's she that allows us to create and manipulate the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myList.Push("Three")
//     fmt.Println(myList.Len())
//     myList.Remove("Two")
//     fmt.Println(myList.Len())
//     myList.Print()
//
type List struct {
	First *Item
}

// Item stores the value and a pointer to the next Item, if the pointer is null, this Item is the last of the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("one")
//     myList.Push("Two")
//     myList.Push("Three")
//     fmt.Println(myList.First.Value)
//     fmt.Println(myList.First.Next.Value)
//     fmt.Println(myList.First.Next.Next.Value)
//     fmt.Println(myList.First.Next.Next.Next)
//
type Item struct {
	Value interface{}
	Next  *Item
}

// Len returns the size of the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myList.Push("Three")
//     fmt.Println(myList.Len())
//
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

// Push a new item into tail of the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//
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

// Find an item in the linked list and return true if it exists or false, if it does not exist. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     fmt.Println(myList.Find("Two"))
//     fmt.Println(myList.Find("Three"))
//
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

// Remove a item of the linked list and return a bool value. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     fmt.Println(myList.Remove("Two"))
//     fmt.Println(myList.Remove("Three"))
//     fmt.Println(myList.Remove("Two"))
//
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

// Print the linked list to stdOut. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myList.Print()
//
func (l *List) Print() {
	p := l.First
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}
}

// Load list from file and push into tail of the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Load("input.txt")
//
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

// Save the linked list to file. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myList.Save("/tmp/myList.txt")
//
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

// Invert the linked list. Example:
//
//     var myList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myList.Inversor()
//     myList.Print()
//
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

// Copy the itens from one linked list to tail of the other linked list. Example:
//
//     var myList linkedlist.List
//     var myOtherList linkedlist.List
//     myList.Push("One")
//     myList.Push("Two")
//     myOtherList.Push("Three")
//     myOtherList.Push("Four")
//     myOtherList.Copy(myList)
//     myOtherList.Print()
//
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
