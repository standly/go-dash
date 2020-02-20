package scte

import "encoding/xml"

type SpliceInfoSection struct {
	ProtocolVersion        *uint64                 `xml:"protocolVersion,attr"`
	PtsAdjustment          *uint64                 `xml:"ptsAdjustment,attr"`
	Tier                   *uint64                 `xml:"tier,attr"`
	TimeSignal             *TimeSignal             `xml:"TimeSignal"`
	SegmentationDescriptor *SegmentationDescriptor `xml:"SegmentationDescriptor"`
}

func (sis *SpliceInfoSection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type SpliceInfoSectionM struct {
		XMLName                xml.Name                `xml:"scte35:SpliceInfoSection"`
		ProtocolVersion        *uint64                 `xml:"protocolVersion,attr"`
		PtsAdjustment          *uint64                 `xml:"ptsAdjustment,attr"`
		Tier                   *uint64                 `xml:"tier,attr"`
		TimeSignal             *TimeSignal             `xml:"scte35:TimeSignal"`
		SegmentationDescriptor *SegmentationDescriptor `xml:"scte35:SegmentationDescriptor"`
	}
	err := e.Encode(&SpliceInfoSectionM{
		ProtocolVersion:        sis.ProtocolVersion,
		PtsAdjustment:          sis.PtsAdjustment,
		Tier:                   sis.Tier,
		TimeSignal:             sis.TimeSignal,
		SegmentationDescriptor: sis.SegmentationDescriptor,
	})
	if err != nil {
		return err
	}
	return nil
}
