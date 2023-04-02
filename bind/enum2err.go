package bind

import "errors"

type Err int

const (
	XmlSuccess Err = iota
	XmlNoAttribute
	XmlWrongAttributeType
	XmlErrorFileNotFound
	XmlErrorFileCouldNotBeOpened
	XmlErrorFileReadError
	XmlErrorParsingElement
	XmlErrorParsingAttribute
	XmlErrorParsingText
	XmlErrorParsingCdata
	XmlErrorParsingComment
	XmlErrorParsingDeclaration
	XmlErrorParsingUnknown
	XmlErrorEmptyDocument
	XmlErrorMismatchedElement
	XmlErrorParsing
	XmlCanNotConvertText
	XmlNoTextNode
	XmlElementDepthExceeded
	XmlErrorCount
)

var errStr = map[Err]error{
	//XmlSuccess:                   errors.New("XML_SUCCESS"),
	XmlNoAttribute:               errors.New("XML_NO_ATTRIBUTE"),
	XmlWrongAttributeType:        errors.New("XML_WRONG_ATTRIBUTE_TYPE"),
	XmlErrorFileNotFound:         errors.New("XML_ERROR_FILE_NOT_FOUND"),
	XmlErrorFileCouldNotBeOpened: errors.New("XML_ERROR_FILE_COULD_NOT_BE_OPENED"),
	XmlErrorFileReadError:        errors.New("XML_ERROR_FILE_READ_ERROR"),
	XmlErrorParsingElement:       errors.New("XML_ERROR_PARSING_ELEMENT"),
	XmlErrorParsingAttribute:     errors.New("XML_ERROR_PARSING_ATTRIBUTE"),
	XmlErrorParsingText:          errors.New("XML_ERROR_PARSING_TEXT"),
	XmlErrorParsingCdata:         errors.New("XML_ERROR_PARSING_CDATA"),
	XmlErrorParsingComment:       errors.New("XML_ERROR_PARSING_COMMENT"),
	XmlErrorParsingDeclaration:   errors.New("XML_ERROR_PARSING_DECLARATION"),
	XmlErrorParsingUnknown:       errors.New("XML_ERROR_PARSING_UNKNOWN"),
	XmlErrorEmptyDocument:        errors.New("XML_ERROR_EMPTY_DOCUMENT"),
	XmlErrorMismatchedElement:    errors.New("XML_ERROR_MISMATCHED_ELEMENT"),
	XmlErrorParsing:              errors.New("XML_ERROR_PARSING"),
	XmlCanNotConvertText:         errors.New("XML_CAN_NOT_CONVERT_TEXT"),
	XmlNoTextNode:                errors.New("XML_NO_TEXT_NODE"),
	XmlElementDepthExceeded:      errors.New("XML_ELEMENT_DEPTH_EXCEEDED"),
	XmlErrorCount:                errors.New("XML_ERROR_COUNT"),
}

func (e Err) Error() error {
	return errStr[e]
}
