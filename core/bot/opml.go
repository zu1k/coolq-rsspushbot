package bot

import (
	"encoding/xml"
)

type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	Title           string `xml:"title"`
	DateCreated     string `xml:"dateCreated,omitempty"`
	DateModified    string `xml:"dateModified,omitempty"`
	OwnerName       string `xml:"ownerName,omitempty"`
	OwnerEmail      string `xml:"ownerEmail,omitempty"`
	OwnerID         string `xml:"ownerId,omitempty"`
	Docs            string `xml:"docs,omitempty"`
	ExpansionState  string `xml:"expansionState,omitempty"`
	VertScrollState string `xml:"vertScrollState,omitempty"`
	WindowTop       string `xml:"windowTop,omitempty"`
	WindowBottom    string `xml:"windowBottom,omitempty"`
	WindowLeft      string `xml:"windowLeft,omitempty"`
	WindowRight     string `xml:"windowRight,omitempty"`
}

type Body struct {
	Outlines []Outline `xml:"outline"`
}

type Outline struct {
	Outlines     []Outline `xml:"outline"`
	Text         string    `xml:"text,attr"`
	Type         string    `xml:"type,attr,omitempty"`
	IsComment    string    `xml:"isComment,attr,omitempty"`
	IsBreakpoint string    `xml:"isBreakpoint,attr,omitempty"`
	Created      string    `xml:"created,attr,omitempty"`
	Category     string    `xml:"category,attr,omitempty"`
	XMLURL       string    `xml:"xmlUrl,attr,omitempty"`
	HTMLURL      string    `xml:"htmlUrl,attr,omitempty"`
	URL          string    `xml:"url,attr,omitempty"`
	Language     string    `xml:"language,attr,omitempty"`
	Title        string    `xml:"title,attr,omitempty"`
	Version      string    `xml:"version,attr,omitempty"`
	Description  string    `xml:"description,attr,omitempty"`
}

func (o OPML) GetFlattenOutlines() ([]Outline, error) {
	var flattenOutlines []Outline
	for _, line := range o.Body.Outlines {
		if line.Outlines != nil {
			for _, subLine := range line.Outlines {
				// 查找子outline
				if subLine.XMLURL != "" {
					flattenOutlines = append(flattenOutlines, subLine)
				}
			}
		}

		if line.XMLURL != "" {
			flattenOutlines = append(flattenOutlines, line)
		}
	}
	return flattenOutlines, nil
}

func (o OPML) XML() (string, error) {
	b, err := xml.MarshalIndent(o, "", "\t")
	return xml.Header + string(b), err
}
