package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// WSDL Types

type WSDLDefinitions struct {
	XMLName      xml.Name       `xml:"definitions"`
	WSDLPortType []WSDLPortType `xml:"portType"`
	WSDLMessage  []WSDLMessage  `xml:"message"`
	WSDLTypes    WSDLTypes      `xml:"types"`
}

type WSDLTypes struct {
	XMLName xml.Name `xml:"types"`
	SSchema SSchema  `xml:"schema"`
}

type SSchema struct {
	XMLName  xml.Name   `xml:"schema"`
	SElement []SElement `xml:"element"`
}

type SElement struct {
	XMLName      xml.Name     `xml:"element"`
	Name         string       `xml:"name,attr"`
	MinOccurs    string       `xml:"minOccurs,attr"`
	MaxOccurs    string       `xml:"maxOccurs,attr"`
	Type         string       `xml:"type,attr"`
	SComplexType SComplexType `xml:"complexType"`
}

type SComplexType struct {
	XMLName   xml.Name  `xml:"complexType"`
	SSequence SSequence `xml:"sequence"`
}

type SSequence struct {
	XMLName  xml.Name   `xml:"sequence"`
	SElement []SElement `xml:"element"`
}

type WSDLPortType struct {
	XMLName    xml.Name        `xml:"portType"`
	Name       string          `xml:"name,attr"`
	Operations []WSDLOperation `xml:"operation"`
}

type WSDLMessage struct {
	XMLName  xml.Name `xml:"message"`
	Name     string   `xml:"name,attr"`
	WSDLPart WSDLPart `xml:"part"`
}

type WSDLPart struct {
	XMLName xml.Name `xml:"part"`
	Name    string   `xml:"name,attr"`
	Element string   `xml:"element,attr"`
}

type WSDLOperation struct {
	XMLName xml.Name `xml:"operation"`
	Name    string   `xml:"name,attr"`
	Input   Input    `xml:"input"`
	Output  Output   `xml:"output"`
}

type Input struct {
	XMLName xml.Name `xml:"input"`
	Message string   `xml:"message,attr"`
}

type Output struct {
	XMLName xml.Name `xml:"output"`
	Message string   `xml:"message,attr"`
}

// Formly types

type Schema struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Type        string              `json:"type"`
	Required    []string            `json:"required"`
	Properties  map[string]Property `json:"properties"`
}

type Property struct {
	// TODO(Peter): Enumerar Type
	Type    string `json:"type"`
	Title   string `json:"title"`
	Default string `json:"default"`
}

func PrintJSON(wsdlByteArray []byte) {
	var w WSDLDefinitions

	xml.Unmarshal(wsdlByteArray, &w)

	for i := 0; i < len(w.WSDLPortType); i++ {
		for j := 0; j < len(w.WSDLPortType[i].Operations); j++ {
			var s = new(Schema)
			var o = w.WSDLPortType[i].Operations[j]

			s.Title = o.Name + " form"
			s.Description = o.Name + " form example."
			s.Type = "object"
			s.Required = []string{}
			s.Properties = make(map[string]Property)

			inputMessageValue := o.Input.Message[4:]

			for m := 0; m < len(w.WSDLMessage); m++ {
				if w.WSDLMessage[m].Name == inputMessageValue {

					messagePartElement := w.WSDLMessage[m].WSDLPart.Element[4:]

					for n := 0; n < len(w.WSDLTypes.SSchema.SElement); n++ {
						if w.WSDLTypes.SSchema.SElement[n].Name == messagePartElement {
							for e := 0; e < len(w.WSDLTypes.SSchema.SElement[n].SComplexType.SSequence.SElement); e++ {
								// TODO(Peter): Como eu sei se o elemento é requerido?
								element := w.WSDLTypes.SSchema.SElement[n].SComplexType.SSequence.SElement[e]

								s.Properties[element.Name] = Property{
									Type:  element.Type[2:],
									Title: element.Name,
								}

							}
						}
					}

					break
				}
			}

			schemaData, _ := json.MarshalIndent(s, "", "    ")
			schemaStr := string(schemaData)

			fmt.Println(schemaStr)
		}
	}

}
