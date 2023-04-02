#include "../tinyxml2/tinyxml2.h"

extern "C" {
  #include "./capi.h"
}

#include <iostream>

struct XMLDocument{
XMLDocument() {
}
tinyxml2::XMLDocument *doc;
};

XMLDocument *NewXMLDocument() {
  std::cout << "OK" << std::endl;
  return new XMLDocument();
}

int XMLDocumentLoadFile(XMLDocument* doc, const char *filename) {
//return doc->doc->LoadFile(filename);
return 1;
}