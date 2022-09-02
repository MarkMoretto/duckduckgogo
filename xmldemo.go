package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

const SampleUrl = `https://duckduckgo.com/?q=BlackBerry&format=xml`

func main() {
	DemoXml()
}

func DemoXml() {
	var ddgr DDGResponse

	bData := retrieveSampleData()

	// Unmarshal the xml using the cpe object
	xml.Unmarshal(bData, &ddgr)

	// prints the Value
	fmt.Printf("XMLName: %#v\n", ddgr.XMLName)
	fmt.Printf("Type: %s\n", ddgr.Type)
	fmt.Printf("Abstract: %s\n", ddgr.Abstract)
	fmt.Printf("Result Ct: %d\n", len(ddgr.Results))
}

func retrieveSampleData() []byte {
	res, err := http.Get(SampleUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}	
