#include "tinyxml2.h"

extern "C" {
  #include "capi.h"
}

struct XMLDocument{
  XMLDocument(bool processEntities, int whitespaceMode) {
    this->doc = new tinyxml2::XMLDocument(processEntities, (tinyxml2::Whitespace)whitespaceMode);
  }
  ~XMLDocument() {
    delete this->doc;
  }
  tinyxml2::XMLDocument *doc;
};

XMLDocument *CNewXMLDocument(bool processEntities, int whitespaceMode) {
  return new XMLDocument(processEntities, whitespaceMode);
}

void CDeleteXMLDocument(XMLDocument *doc) {
  delete doc;
}

int CXMLDocumentLoadFile(XMLDocument* doc, const char *filename) {
  return doc->doc->LoadFile(filename);
}

int CXMLDocumentParse(XMLDocument *doc, const char *xml, size_t nBytes) {
  return doc->doc->Parse(xml, nBytes);
}