package main

// https://duckduckgo.com/api

import (
	"encoding/xml"
	"fmt"
)

type A struct {
	XMLName xml.Name `xml:"a"`
	Href    string   `xml:"href,attr"`
}

type Icon struct {
	XMLName xml.Name `xml:"Icon"`
	Url     string   `xml:",chardata"`
	Height  int      `xml:"height,attr"`
	Width   int      `xml:"width,attr"`
}

type Result struct {
	XMLName xml.Name `xml:"Result"`
	Value   string   `xml:",chardata"`
	A       A        `xml:"a"`
	Icon
	Text     string
	FirstURL string
}

type RelatedTopic struct {
	XMLName xml.Name `xml:"RelatedTopic"`
	Value   string   `xml:",chardata"`
	A       A        `xml:"a"`
	Icon
	Text     string
	FirstURL string
}

type RelatedTopicsSection struct {
	XMLName xml.Name `xml:"RelatedTopicsSection"`
	Name 	string   `xml:"name,attr"`
}

// Main DuckDuckGo response xml struct.
type XmlResponse struct {
	XMLName              xml.Name `xml:"DuckDuckGoResponse"`
	Version              string   `xml:"version,attr"`
	Type                 string
	Redirect             string
	Heading              string
	Image                string
	ImageWidth			 int
	ImageHeight			 int
	ImageLogo			 int
	Abstract             string
	AbstractText         string
	AbstractURL          string
	AbstractSource       string	
	Infobox				 string
	Entity				 string
	Answer               string
	AnswerType           string
	Definition           string
	DefinitionSource     string
	DefinitionURL        string
	Results              []Result		`xml:"Results>Result"`
	RelatedTopics        []RelatedTopic `xml:"RelatedTopics>RelatedTopic"`
	RelatedTopicsSection []RelatedTopic `xml:"RelatedTopicsSection>RelatedTopic"`
	
}

// ---- Formatting and functions ---

const IconFormat string = `%s
	Url:    %s
	Height: %d
	Width:  %d
`

func (ic Icon) String() string {
	return fmt.Sprintf(IconFormat, "Icon", ic.Url, ic.Height, ic.Width)
}

const ResultFormat string = `%s
	Value:     %s
	Icon:      %v
	Text:  	   %s
	First URL: %s
`

func (r Result) String() string {
	return fmt.Sprintf(ResultFormat, "Result", TrimWS(r.Value),  r.Icon, r.Text, r.FirstURL)
}

func (rt RelatedTopic) String() string {
	return fmt.Sprintf(ResultFormat, "RelatedTopic", TrimWS(rt.Value), rt.Icon, rt.Text, rt.FirstURL)
}
