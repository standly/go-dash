package scte

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
}
