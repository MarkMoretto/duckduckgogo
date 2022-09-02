package main

import "encoding/xml"

// https://duckduckgo.com/api

// https://marduc812.com/2020/12/17/complete-xml-parsing-guide-with-golang/

const SampleQuery string = `https://duckduckgo.com/?q=blackberry&format=xml`

/*
<DuckDuckGoResponse version="1.0">
	<Type>D</Type>
	<Redirect/>
	<Heading>Blackberry</Heading>
	<Image/>
	<ImageWidth>0</ImageWidth>
	<ImageHeight>0</ImageHeight>
	<ImageIsLogo>0</ImageIsLogo>
	<Abstract/>
	<AbstractText/>
	<AbstractURL> </AbstractURL>
	<AbstractSource> </AbstractSource>
	<Infobox/>
	<Entity/>
	<Definition/>
	<Results>
	<Result>
	<a href="https://duckduckgo.com"><b>Official site</b></a><a href="https://duckduckgo.com"></a>
	<Icon height="16" width="16">/i/duckduckgo.com.ico</Icon>
	<Text>Official site</Text>
	<FirstURL>https://duckduckgo.com</FirstURL>
	</Result>
	</Results>
	<RelatedTopics>
		<RelatedTopic>
			<a href=""></a>
			<Icon height="00" width="00"> </Icon>
			<Text> </Text>
			<FirstURL></FirstURL>
		</RelatedTopic>
		<RelatedTopicsSection name="See also">
			<RelatedTopic>
				<a href=""></a>
				<Icon> </Icon>
				<Text> </Text>
				<FirstURL></FirstURL>
			</RelatedTopic>
		</RelatedTopicsSection>
	</RelatedTopics>
</DuckDuckGoResponse>
*/

// ----------------------------------------------------------------
// ----------------------------------------------------------------

/*
Return fields -
Abstract: topic summary (can contain HTML, e.g. italics)
AbstractText: topic summary (with no HTML)
AbstractSource: name of Abstract source
AbstractURL: deep link to expanded topic page in AbstractSource
Image: link to image that goes with Abstract
Heading: name of topic that goes with Abstract

Answer: instant answer
AnswerType: type of Answer, e.g. calc, color, digest, info, ip, iploc, phone, pw, rand, regexp, unicode, upc, or zip (see the tour page for examples).

Definition: dictionary definition (may differ from Abstract)
DefinitionSource: name of Definition source
DefinitionURL: deep link to expanded definition page in DefinitionSource

RelatedTopics: array of internal links to related topics associated with Abstract
  Result: HTML link(s) to related topic(s)
  FirstURL: first URL in Result
  Icon: icon associated with related topic(s)
    URL: URL of icon
    Height: height of icon (px)
    Width: width of icon (px)
  Text: text from first URL

Results: array of external links associated with Abstract
  Result: HTML link(s) to external site(s)
  FirstURL: first URL in Result
  Icon: icon associated with FirstURL
    URL: URL of icon
    Height: height of icon (px)
    Width: width of icon (px)
  Text: text from FirstURL

Type: response category, i.e. A (article), D (disambiguation), C (category), N (name), E (exclusive), or nothing.

Redirect: !bang redirect URL
*/

type Icon struct {
	XMLName xml.Name `xml:"Icon"`
	Url string `xml:",chardata"`
	Height int `xml:"height,attr"`
	Width int `xml:"width,attr"`
}

type Result struct {
	XMLName xml.Name `xml:"Result"`
	Value string `xml:",chardata"`
	Icon
	Text string
	FirstURL string
}

type RelatedTopic struct {
	XMLName xml.Name `xml:"RelatedTopic"`
	Value string `xml:",chardata"`
	Icon
	Text string
	FirstURL string
}


type RelatedTopicsSection struct {
	XMLName xml.Name `xml:"RelatedTopicsSection"`
	Name string `xml:"name,attr"`
}

//TODO: Refine this.
type DDGResponse struct {
	XMLName xml.Name `xml:"DuckDuckGoResponse"`
	Version string `xml:"version,attr"`
	Abstract string
	AbstractText string
	AbstractSource string
	AbstractURL string
	Image string
	Heading string
	Answer string
	Redirect string
	AnswerType string
	Definition string
	DefinitionSource string
	DefinitionURL string
	RelatedTopics []RelatedTopic
	Results []Result
	Type string
}
