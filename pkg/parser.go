package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
)


func ReadLocalFile(filePath ...string) ([]byte, error) {
	var f *os.File
	var err error
	var bData []byte
	// Handle file path argument.
	if len(filePath) == 0 || filePath == nil {
		return nil, fmt.Errorf("Error: Missing parameter `filePath`")
	}

	// Read file and check for error.
	f, err = os.Open(filePath[0])
	if err != nil {
		return nil, fmt.Errorf("Error opening and reading file")
	}
	defer f.Close()

	bData, _ = ioutil.ReadAll(f)
	return bData, nil
}

// func ParseXmlObject(xmlObj struct{}) {
// 	var fileData []byte
// 	var err error	
// 	if fileData, err = ReadLocalFile(XmlFilePath); err != nil {
// 		log.Fatal(err)
// 	}
// 	switch {
// 	case xmlObj.(Type) == DDGResponse:
// 	}
// 	var ddgr DDGResponse


// }