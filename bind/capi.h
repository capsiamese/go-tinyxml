#pragma once

typedef struct XMLDocument XMLDocument;

XMLDocument* NewXMLDocument();
int XMLDocumentLoadFile(XMLDocument *doc, const char* filename);