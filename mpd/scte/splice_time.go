package scte

import "encoding/xml"

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
