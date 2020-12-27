package core

import (
	"encoding/xml"
	"fmt"
)

type WSDLDefinitions struct {
	XMLName     xml.Name      `xml:"definitions"`
	WSDLBinding []WSDLBinding `xml:"binding"`
}

type WSDLBinding struct {
	XMLName    xml.Name        `xml:"binding"`
	Name       string          `xml:"name,attr"`
	Type       string          `xml:"type,attr"`
	Operations []WSDLOperation `xml:"operation"`
}

type WSDLOperation struct {
	XMLName xml.Name `xml:"operation"`
	Name    string   `xml:"name,attr"`
}

func PrintJSON(wsdlByteArray []byte) {
	var w WSDLDefinitions

	xml.Unmarshal(wsdlByteArray, &w)

	fmt.Println(w.WSDLBinding[0].Name)
	for i := 0; i < len(w.WSDLBinding); i++ {
		fmt.Println("Binding Name: " + w.WSDLBinding[i].Name)
		for j := 0; j < len(w.WSDLBinding[i].Operations); j++ {
			fmt.Println("Operation Name: " + w.WSDLBinding[i].Operations[j].Name)
		}
	}
}
