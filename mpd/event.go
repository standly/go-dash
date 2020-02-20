package mpd

import (
	"encoding/xml"
	"github.com/standly/go-dash/mpd/scte"
)

type EventBase struct {
	Timescale        *int64 `xml:"timescale,attr,omitempty"`
	PresentationTime *int64 `xml:"presentationTime,attr,omitempty"`
	Duration         *int64 `xml:"duration,attr,omitempty"`
	ID               *int64 `xml:"id,attr,omitempty"`
}

type Event struct {
	EventBase
	// https://stackoverflow.com/questions/34820549/unable-to-parse-xml-in-go-with-in-tags
	// wired is `xml:"scte35 SpliceInfoSection,omitempty"` not working here
	SpliceInfoSection *scte.SpliceInfoSection `xml:"SpliceInfoSection"`
}

type EventStream struct {
	SchemeIDUri *string  `xml:"schemeIdUri,attr"`
	Timescale   *int64   `xml:"timescale,attr"`
	Event       []*Event `xml:"Event,omitempty"`
}

type EventMarshal struct {
	XMLName xml.Name `xml:"Event"`
	EventBase
	// https://stackoverflow.com/questions/34820549/unable-to-parse-xml-in-go-with-in-tags
	// wired is `xml:"scte35 SpliceInfoSection,omitempty"` not working here
	SpliceInfoSection *scte.SpliceInfoSection `xml:"scte35:SpliceInfoSection"`
}

func (event *Event) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.Encode(&EventMarshal{
		EventBase:         event.EventBase,
		SpliceInfoSection: event.SpliceInfoSection,
	})
	if err != nil {
		return err
	}
	return nil
}
