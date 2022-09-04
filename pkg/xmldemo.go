package pkg

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

const SampleUrl = `https://duckduckgo.com/?q=blackberry&format=xml`
// const SampleUrl = `https://duckduckgo.com/?q=BlackBerry&format=xml`

func main() {
	DemoXmlLocalFile()
	// DemoXml()
}

// --- Local XML file demo. ---

const XmlFilePath string = `data/sample.xml`

func DemoXmlLocalFile() {
	var fileData []byte
	var err error	
	if fileData, err = ReadLocalFile(XmlFilePath); err != nil {
		log.Fatal(err)
	}

	ddgr := &XmlResponse{}
	err = xml.Unmarshal(fileData, &ddgr)
	if err != nil {
		log.Fatal(err)
	}

	for _, rt := range ddgr.RelatedTopics {
		fmt.Printf("%v\n", rt)
	}

}

// --- Internet-based XML data demo. ---
func DemoXml() {
	ddgr := &XmlResponse{}

	bData := retrieveSampleData()

	// Unmarshal the xml using the cpe object
	xml.Unmarshal(bData, &ddgr)

	// prints the Value
	fmt.Printf("XMLName: %#v\n", ddgr.XMLName)
	fmt.Printf("Type: %s\n", ddgr.Type)
	fmt.Printf("Version: %s\n", ddgr.Version)
	fmt.Printf("Abstract: %s\n", ddgr.Abstract)
	fmt.Printf("AbstractText: %s\n", ddgr.AbstractText)
	fmt.Printf("AbstractSource: %s\n", ddgr.AbstractSource)
	fmt.Printf("AbstractURL: %s\n", ddgr.AbstractURL)
	fmt.Printf("Image: %s\n", ddgr.Image)
	fmt.Printf("Heading: %s\n", ddgr.Heading)
	fmt.Printf("Answer: %s\n", ddgr.Answer)
	fmt.Printf("Redirect: %s\n", ddgr.Redirect)
	fmt.Printf("AnswerType: %s\n", ddgr.AnswerType)
	fmt.Printf("Definition: %s\n", ddgr.Definition)
	fmt.Printf("DefinitionSource: %s\n", ddgr.DefinitionSource)
	fmt.Printf("DefinitionURL: %s\n", ddgr.DefinitionURL)

	fmt.Printf("Results Count: %d\n", len(ddgr.Results))
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
