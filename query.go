package main

// Ref: https://duckduckgo.com/api

/*
Query parameters

q: query
format: output format (json or xml)
	If format=='json', you can also pass:
		- callback: function to callback (JSONP format)

pretty: 1 to make JSON look pretty (like JSONView for Chrome/Firefox)
no_redirect: 1 to skip HTTP redirects (for !bang commands).
no_html: 1 to remove HTML from text, e.g. bold and italics.
skip_disambig: 1 to skip disambiguation (D) Type.
*/

// https://duckduckgo.com/?q=blackberry&format=xml
// https://duckduckgo.com/?q=valley+forge+national+park&format=xml
// https://duckduckgo.com/?q=simpsons+characters&format=xml

import (
	"fmt"
	"strings"
)



const baseQuery string = `https://duckduckgo.com`

func GetBaseQuery() string {
	return baseQuery
}

func generateUrlString(qs string, rf ReturnFormat) string {
	return fmt.Sprintf("%s/?format=xml", baseQuery)
}

// Create url with given query from given parameter.
// Formatted to return 
func XmlQuery(queryString string) string {
	return fmt.Sprintf("%s/?q=%s&format=xml", baseQuery, Querify(queryString))
}

func Querify(qs string) string {
	qparts := strings.Fields(qs)
	return strings.Join(qparts, "+")
}

type SimpleTestCase map[string]string

func DemoXmlQuery() {
	var testCases SimpleTestCase
	testCases = SimpleTestCase{
		"blackberry": "https://duckduckgo.com/?q=blackberry&format=xml",
		"valley forge national park": "https://duckduckgo.com/?q=valley+forge+national+park&format=xml",
	}
	
	for queryArg, givenUrl := range testCases {
		fmt.Printf(formatTC(XmlQuery(queryArg), givenUrl))
	}
}

// Format test case output.
const TestResultFormat = `
Actual:   %s
Expected: %s
`

func formatTC(a, e string) string {
	return strings.TrimLeft(fmt.Sprintf(TestResultFormat, a, e), " ")
}
