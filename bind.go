package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -std=c11
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
	//C.CString("")
	//C.GOString(nil)
	res := C.CXMLDocumentLoadFile((*C.XMLDocument)(c.doc), C.CString(fileName))
	return (Err)(res)
}

func (c *cXMLDocumentWrap) Parse(xml []byte) Err {
	return (Err)(C.CXMLDocumentParse((*C.XMLDocument)(c.doc),
		C.CString(string(xml)), C.size_t(len(xml)),
	))
}

func NewXMLDocument() XMLDocument {
	var processEntities = true
	var whitespaceMode = 0
	return &cXMLDocumentWrap{
		doc: (*cXMLDocument)(C.CNewXMLDocument(C.bool(processEntities), C.int(whitespaceMode))),
	}
}

type XMLDocument interface {
	Todo()
	LoadFile(fileName string) Err
	Parse(xml []byte) Err
}
