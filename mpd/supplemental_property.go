package mpd

import "time"

type SupplementalProperty struct {
	SchemeIdUri *string `xml:"schemeIdUri,attr,omitempty"`
	Value       *string `xml:"value,attr,omitempty"`
}

const (
	ScteUTCTimeSchemeIDURI = "urn:scte:dash:utc-time"
)

func NewSupplementalPropertyWithScteUTCTime() *SupplementalProperty {
	t := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	s := ScteUTCTimeSchemeIDURI
	return &SupplementalProperty{
		SchemeIdUri: &s,
		Value:       &t,
	}
}
