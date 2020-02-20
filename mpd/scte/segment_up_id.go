package scte

import (
	"encoding/xml"
	"strings"
)

type SegmentationUpID struct {
	SegmentationUpidType   *uint64 `xml:"segmentationUpidType,attr"`
	SegmentationUpidLength *uint64 `xml:"segmentationUpidLength,attr"`
	SegmentationTypeId     *uint64 `xml:"segmentationTypeId,attr"`
	SegmentNum             *uint64 `xml:"segmentNum,attr"`
	SegmentsExpected       *uint64 `xml:"segmentsExpected,attr"`
	SegmentationUpidFromat *string `xml:"segmentationUpidFromat,attr"`
	SegmentationUpIDValue  *string `xml:",chardata"`
}

func (sui *SegmentationUpID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if sui.SegmentationUpIDValue != nil {
		idv := strings.TrimSpace(*sui.SegmentationUpIDValue)
		sui.SegmentationUpIDValue = &idv
	}
	type SegmentationUpIDMashal struct {
		XMLName                xml.Name `xml:"scte35:SegmentationUpid"`
		SegmentationUpidType   *uint64  `xml:"segmentationUpidType,attr"`
		SegmentationUpidLength *uint64  `xml:"segmentationUpidLength,attr"`
		SegmentationTypeId     *uint64  `xml:"segmentationTypeId,attr"`
		SegmentNum             *uint64  `xml:"segmentNum,attr"`
		SegmentsExpected       *uint64  `xml:"segmentsExpected,attr"`
		SegmentationUpidFromat *string  `xml:"segmentationUpidFromat,attr"`
		SegmentationUpIDValue  *string  `xml:",chardata"`
	}
	err := e.Encode(&SegmentationUpIDMashal{
		SegmentationUpidType:   sui.SegmentationUpidType,
		SegmentationUpidLength: sui.SegmentationUpidLength,
		SegmentationTypeId:     sui.SegmentationTypeId,
		SegmentNum:             sui.SegmentNum,
		SegmentsExpected:       sui.SegmentsExpected,
		SegmentationUpIDValue:  sui.SegmentationUpIDValue,
		SegmentationUpidFromat: sui.SegmentationUpidFromat,
	})
	if err != nil {
		return err
	}
	return nil
}
