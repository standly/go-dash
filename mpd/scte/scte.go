package scte

const (
	DASH_SCTE35_URI = "urn:scte:scte35:2013:xml"
)

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

type DeliveryRestrictions struct {
	WebDeliveryAllowedFlag *bool `xml:"webDeliveryAllowedFlag,attr"`
	NoRegionalBlackoutFlag *bool `xml:"noRegionalBlackoutFlag,attr"`
	ArchiveAllowedFlag     *bool `xml:"archiveAllowedFlag,attr"`
	DeviceRestrictions     int64 `xml:"deviceRestrictions,attr"`
}
