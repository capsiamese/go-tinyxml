package bind

/*
#cgo CXXFLAGS: -std=c++11
#include "capi.h"
*/
import "C"
import "fmt"

type cXMLDocument C.XMLDocument

type cXMLDocumentWrap struct {
	doc *cXMLDocument
}

func (c *cXMLDocumentWrap) Todo() {
	fmt.Println("TODO")
}

func (c *cXMLDocumentWrap) LoadFile(fileName string) Err {
	res := C.XMLDocumentLoadFile((*C.XMLDocument)(c.doc), C.CString(fileName))
	return (Err)(res)
}

func NewXMLDocument() XMLDocument {
	return &cXMLDocumentWrap{
		doc: (*cXMLDocument)(C.NewXMLDocument()),
	}
}

type XMLDocument interface {
	Todo()
	LoadFile(fileName string) Err
}
