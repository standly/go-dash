package scte

import (
	"encoding/xml"
	"strings"
)

const (
	DASH_SCTE35_URI = "urn:scte:scte35:2013:xml"
)

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

type SpliceInsert struct {
	SpliceEventId              uint64 `xml:"spliceEventId,attr"`
	SpliceEventCancelIndicator bool   `xml:"spliceEventCancelIndicator,attr"`
	OutOfNetworkIndicator      bool   `xml:"outOfNetworkIndicator,attr"`
	UniqueProgramId            uint64 `xml:"uniqueProgramId,attr"`
	AvailNum                   int64  `xml:"availNum,attr"`
	AvailsExpected             int64  `xml:"availsExpected,attr"`
	SpliceImmediateFlag        bool   `xml:"spliceImmediateFlag,attr"`
}

type Program struct {
	SpliceTime *SpliceTime `xml:"SpliceTime"`
}

type SpliceTime struct {
	PtsTime *uint64 `xml:"ptsTime,attr"`
}

func (sd *SpliceTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type SpliceTimeM struct {
		XMLName xml.Name `xml:"scte35:SpliceTime"`
		PtsTime *uint64  `xml:"ptsTime,attr"`
	}
	err := e.Encode(&SpliceTimeM{
		PtsTime: sd.PtsTime,
	})
	if err != nil {
		return err
	}
	return nil
}

type BreakDuration struct {
	AutoReturn *bool   `xml:"autoReturn,attr"`
	Duration   *uint64 `xml:"duration,attr"`
}

type AvailDescriptor struct {
	ProviderAvailId *uint64 `xml:"providerAvailId,attr"`
}

type TimeSignal struct {
	SpliceTime *SpliceTime `xml:"SpliceTime"`
}

type SegmentationDescriptor struct {
	SegmentationEventId              *uint64               `xml:"segmentationEventId,attr"`
	SegmentationEventCancelIndicator *bool                 `xml:"segmentationEventCancelIndicator,attr"`
	DeliveryRestrictions             *DeliveryRestrictions `xml:"DeliveryRestrictions"`
	SegmentationUpID                 *SegmentationUpID     `xml:"SegmentationUpid"`
}

func (sd *SegmentationDescriptor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type SegmentationDescriptorMashal struct {
		XMLName                          xml.Name              `xml:"scte35:SegmentationDescriptor"`
		SegmentationEventId              *uint64               `xml:"segmentationEventId,attr"`
		SegmentationEventCancelIndicator *bool                 `xml:"segmentationEventCancelIndicator,attr"`
		DeliveryRestrictions             *DeliveryRestrictions `xml:"scte35:DeliveryRestrictions"`
		SegmentationUpID                 *SegmentationUpID     `xml:"scte35:SegmentationUpid"`
	}
	err := e.Encode(&SegmentationDescriptorMashal{
		SegmentationEventId:              sd.SegmentationEventId,
		SegmentationEventCancelIndicator: sd.SegmentationEventCancelIndicator,
		DeliveryRestrictions:             sd.DeliveryRestrictions,
		SegmentationUpID:                 sd.SegmentationUpID,
	})
	if err != nil {
		return err
	}
	return nil
}

type DeliveryRestrictions struct {
	WebDeliveryAllowedFlag *bool `xml:"webDeliveryAllowedFlag,attr"`
	NoRegionalBlackoutFlag *bool `xml:"noRegionalBlackoutFlag,attr"`
	ArchiveAllowedFlag     *bool `xml:"archiveAllowedFlag,attr"`
	DeviceRestrictions     int64 `xml:"deviceRestrictions,attr"`
}

type SegmentationUpID struct {
	SegmentationUpidType   *uint64 `xml:"segmentationUpidType,attr"`
	SegmentationUpidLength *uint64 `xml:"segmentationUpidLength,attr"`
	SegmentationTypeId     *uint64 `xml:"segmentationTypeId,attr"`
	SegmentNum             *uint64 `xml:"segmentNum,attr"`
	SegmentsExpected       *uint64 `xml:"segmentsExpected,attr"`
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
		SegmentationUpIDValue  *string  `xml:",chardata"`
	}
	err := e.Encode(&SegmentationUpIDMashal{
		SegmentationUpidType:   sui.SegmentationUpidType,
		SegmentationUpidLength: sui.SegmentationUpidLength,
		SegmentationTypeId:     sui.SegmentationTypeId,
		SegmentNum:             sui.SegmentNum,
		SegmentsExpected:       sui.SegmentsExpected,
		SegmentationUpIDValue:  sui.SegmentationUpIDValue,
	})
	if err != nil {
		return err
	}
	return nil
}
