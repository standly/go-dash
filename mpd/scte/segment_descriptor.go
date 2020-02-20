package scte

import "encoding/xml"

type SegmentationDescriptor struct {
	SegmentationEventId              *uint64 `xml:"segmentationEventId,attr"`
	SegmentationEventCancelIndicator *bool   `xml:"segmentationEventCancelIndicator,attr"`

	SegmentationDuration *uint64 `xml:"segmentationDuration,attr"`
	SegmentationTypeId   *uint64 `xml:"segmentationTypeId,attr"`
	SegmentNum           *uint64 `xml:"segmentNum,attr"`
	SegmentsExpected     *uint64 `xml:"segmentsExpected,attr"`
	SubSegmentNum        *uint64 `xml:"subSegmentNum,attr"`
	SubSegmentsExpected  *uint64 `xml:"subSegmentsExpected,attr"`

	DeliveryRestrictions *DeliveryRestrictions `xml:"DeliveryRestrictions"`
	SegmentationUpID     *SegmentationUpID     `xml:"SegmentationUpid"`
}

func (sd *SegmentationDescriptor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type SegmentationDescriptorMashal struct {
		XMLName                          xml.Name              `xml:"scte35:SegmentationDescriptor"`
		SegmentationEventId              *uint64               `xml:"segmentationEventId,attr"`
		SegmentationEventCancelIndicator *bool                 `xml:"segmentationEventCancelIndicator,attr"`
		DeliveryRestrictions             *DeliveryRestrictions `xml:"scte35:DeliveryRestrictions"`
		SegmentationUpID                 *SegmentationUpID     `xml:"scte35:SegmentationUpid"`

		SegmentationDuration *uint64 `xml:"segmentationDuration,attr"`
		SegmentationTypeId   *uint64 `xml:"segmentationTypeId,attr"`
		SegmentNum           *uint64 `xml:"segmentNum,attr"`
		SegmentsExpected     *uint64 `xml:"segmentsExpected,attr"`
		SubSegmentNum        *uint64 `xml:"subSegmentNum,attr"`
		SubSegmentsExpected  *uint64 `xml:"subSegmentsExpected,attr"`
	}
	err := e.Encode(&SegmentationDescriptorMashal{
		SegmentationEventId:              sd.SegmentationEventId,
		SegmentationEventCancelIndicator: sd.SegmentationEventCancelIndicator,
		DeliveryRestrictions:             sd.DeliveryRestrictions,
		SegmentationUpID:                 sd.SegmentationUpID,

		SegmentationDuration: sd.SegmentationDuration,
		SegmentationTypeId:   sd.SegmentationTypeId,
		SegmentNum:           sd.SegmentNum,
		SegmentsExpected:     sd.SegmentsExpected,
		SubSegmentNum:        sd.SubSegmentNum,
		SubSegmentsExpected:  sd.SubSegmentsExpected,
	})
	if err != nil {
		return err
	}
	return nil
}
