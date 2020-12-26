package wsdl

import (
	"testing"
)

func TestGetWSDL(t *testing.T) {
	uri := "http://www.dneonline.com/calculator.asmx?WSDL"
	wsdl, _ := GetWSDL(uri)
	if len(wsdl) <= 0 {
		t.Errorf("There was a problem when trying to display wsdl from this uri: '" + uri + "'.")
	}
}
