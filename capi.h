#pragma once

#include <stdbool.h>

typedef struct XMLDocument XMLDocument;

XMLDocument* CNewXMLDocument(bool processEntities, int whitespaceMode);
void CDeleteXMLDocument(XMLDocument *doc);

int CXMLDocumentLoadFile(XMLDocument *doc, const char* filename);
int CXMLDocumentParse(XMLDocument *doc, const char *xml, size_t nBytes);