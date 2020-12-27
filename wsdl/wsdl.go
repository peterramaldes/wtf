package wsdl

import (
	"io/ioutil"
	"net/http"
)

func GetWSDL(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	wsdl, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return wsdl, err
}
