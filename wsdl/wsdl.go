package wsdl

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetWSDL(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, errors.New("Could not get wsdl '" + uri + "'.")
	}

	wsdl, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, errors.New("Could not get wsdl '" + uri + "'.")
	}

	return wsdl, err
}
