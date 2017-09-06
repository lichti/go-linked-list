# linkedlist
--
    import "github.com/lichti/go-linked-list"


## Usage

#### type Item

```go
type Item struct {
	Value interface{}
	Next  *Item
}
```

Item stores the value and a pointer to the next Item, if the pointer is null,
this Item is the last of the linked list. Example:

    var myList linkedlist.List
    myList.Push("one")
    myList.Push("Two")
    myList.Push("Three")
    fmt.Println(myList.First.Value)
    fmt.Println(myList.First.Next.Value)
    fmt.Println(myList.First.Next.Next.Value)
    fmt.Println(myList.First.Next.Next.Next)

#### type List

```go
type List struct {
	First *Item
}
```

List is the main object this library, it's she that allows us to create and
manipulate the linked list. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myList.Push("Three")
    fmt.Println(myList.Len())
    myList.Remove("Two")
    fmt.Println(myList.Len())
    myList.Print()

#### func (*List) Copy

```go
func (l *List) Copy(origin List)
```
Copy the itens from one linked list to tail of the other linked list. Example:

    var myList linkedlist.List
    var myOtherList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myOtherList.Push("Three")
    myOtherList.Push("Four")
    myOtherList.Copy(myList)
    myOtherList.Print()

#### func (*List) Find

```go
func (l *List) Find(value interface{}) bool
```
Find an item in the linked list and return true if it exists or false, if it
does not exist. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    fmt.Println(myList.Find("Two"))
    fmt.Println(myList.Find("Three"))

#### func (*List) Inversor

```go
func (l *List) Inversor()
```
Invert the linked list. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myList.Inversor()
    myList.Print()

#### func (*List) Len

```go
func (l *List) Len() int
```
Len returns the size of the linked list. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myList.Push("Three")
    fmt.Println(myList.Len())

#### func (*List) Load

```go
func (l *List) Load(path string) error
```
Load list from file and push into tail of the linked list. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Load("input.txt")

#### func (*List) Print

```go
func (l *List) Print()
```
Print the linked list to stdOut. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myList.Print()

#### func (*List) Push

```go
func (l *List) Push(value interface{})
```
Push a new item into tail of the linked list. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")

#### func (*List) Remove

```go
func (l *List) Remove(value interface{}) bool
```
Remove a item of the linked list and return a bool value. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    fmt.Println(myList.Remove("Two"))
    fmt.Println(myList.Remove("Three"))
    fmt.Println(myList.Remove("Two"))

#### func (*List) Save

```go
func (l *List) Save(path string) error
```
Save the linked list to file. Example:

    var myList linkedlist.List
    myList.Push("One")
    myList.Push("Two")
    myList.Save("/tmp/myList.txt")
