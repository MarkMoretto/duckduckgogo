package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

const testFilePath string = `./data/sample.xml`

func CheckXmlSchema() {
	f, err := os.Open(testFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	qData, _ := ioutil.ReadAll(f)
	var ddgResp DDGResponse
	xml.Unmarshal(qData, &ddgResp)
}