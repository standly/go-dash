package mpd

import "github.com/standly/go-dash/mpd/scte"

type Event struct {
	Timescale        *int64 `xml:"timescale,attr,omitempty"`
	PresentationTime *int64 `xml:"presentationTime,attr,omitempty"`
	Duration         *int64 `xml:"duration,attr,omitempty"`
	ID               *int64 `xml:"id,attr,omitempty"`
	// https://stackoverflow.com/questions/34820549/unable-to-parse-xml-in-go-with-in-tags
	// wired is `xml:"scte35 SpliceInfoSection,omitempty"` not working here
	SpliceInfoSection *scte.SpliceInfoSection `xml:"scte35\:SpliceInfoSection"`
}

type EventStream struct {
	SchemeIDUri *string  `xml:"schemeIdUri,attr"`
	Timescale   *int64   `xml:"timescale,attr"`
	Event       []*Event `xml:"Event,omitempty"`
}

//func (es EventStream)()  {
//
//}
