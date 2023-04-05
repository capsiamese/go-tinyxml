package main

import (
	"fmt"
)

func main() {
	doc := NewXMLDocument()
	doc.Todo()
	if err := doc.LoadFile("Hello").Error(); err != nil {
		fmt.Println(err)
	}
	if err := doc.Parse([]byte{1, 2, 3}).Error(); err != nil {
		fmt.Println(err)
	}
}
