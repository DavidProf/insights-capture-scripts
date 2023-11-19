package bkpmanager

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func GetDataFromFileAsBytes(filepath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(string(fileBytes[:]))
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	fz, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	finalContentBytes, err := ioutil.ReadAll(fz)
	if err != nil {
		return nil, err
	}

	return finalContentBytes, nil
}

func GetDataFromFileAsMap(filepath string) (map[string]any, error) {
	finalContentBytes, err := GetDataFromFileAsBytes(filepath)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	json.Unmarshal([]byte(finalContentBytes), &result)

	return result, nil
}
