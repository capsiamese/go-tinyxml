package main

import (
	"fmt"
	"github.com/capsiamese/go-tinyxml/bind"
)

func main() {
	doc := bind.NewXMLDocument()
	doc.Todo()
	if err := doc.LoadFile("Hello").Error(); err != nil {
		fmt.Println(err)
	}
}
