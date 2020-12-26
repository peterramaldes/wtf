package wsdl

import (
	"testing"
)

func TestGetWSDL(t *testing.T) {
	uri := "http://www.dneonline.com/calculator.asmx?WSDL"
	wsdl, _ := GetWSDL(uri)
	if len(wsdl) <= 0 {
		t.Errorf("Cannot get wsdl in this uri: '" + uri + "'.")
	}
}
