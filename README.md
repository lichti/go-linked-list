# go-linked-list
Go lib to create and manipule linked lists

# Simple example
```
package main

import (
	"github.com/lichti/go-linked-list"
)

func main() {
	var input linkedlist.List
	var output linkedlist.List
	input.Load("entrada.txt")
	output.Copy(input)
	output.Inversor()
	output.Save("saida.txt")
}
```

# License
GPLv3

# Author
 - [Gustavo Lichti](gustavo.lichti@gmail.com)