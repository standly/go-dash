package main

import (
	"bytes"
	"fmt"
	"github.com/standly/go-dash/mpd"
	"io/ioutil"
)

func main() {
	xmlContent, e := ioutil.ReadFile("examples_scte35/scte35_examples.xml")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	playlist, e := mpd.Read(bytes.NewReader(xmlContent))
	if e != nil {
		fmt.Println(e)
		return
	}
	es := playlist.Periods[0].EventStream
	fmt.Println(es.Event[0].ID)
}
