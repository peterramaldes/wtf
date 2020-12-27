package wsdl

import (
	"testing"
)

func TestGetWSDL(t *testing.T) {
	uri := "http://www.dneonline.com/calculator.asmx?WSDL"
	wsdl, err := GetWSDL(uri)

	if len(wsdl) <= 0 {
		t.Error(err)
	}
}
