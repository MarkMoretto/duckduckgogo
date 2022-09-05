package pkg

// Enum tests.

import (
	"fmt"
	"testing"
)

func TestJsonOkay(t testing.T) {
	res := fmt.Sprintf("%s\n", Json)
	if res != "json" {
		t.Errorf("Returned %s expected `json`", res)
	}
}

func TestXmlOkay(t testing.T) {
	res := fmt.Sprintf("%s\n", Xml)
	if res != "xml" {
		t.Errorf("Returned %s expected `xml`", res)
	}
}