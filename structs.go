package main

import (
	"encoding/xml"

	"github.com/simplereach/timeutils"
)

// Channel : channel struct
type Channel struct {
	XMLName     xml.Name       `xml:"channel"`
	ChID        string         `json:"ch_id" xml:"id,attr"`
	ChannelName string         `json:"channel_name" xml:"display-name"`
	Rec         string         `json:"rec" xml:"-"`
	Img         string         `json:"img" xml:"-"`
	Imgsrc      Img            `xml:"icon"`
	Category    Category       `json:"category" xml:"-"`
	Name        string         `json:"name" xml:"-"`
	Time        timeutils.Time `json:"time" xml:"-"`
	TimeTo      timeutils.Time `json:"time_to" xml:"-"`
	Duration    string         `json:"duration" xml:"-"`
	Descr       string         `json:"descr" xml:"-"`
	Programm    []Programm
}

// Img : channel struct
type Img struct {
	Img string `json:"" xml:"src,attr"`
}

// Programm : channel struct
type Programm struct {
	Rec      string         `json:"rec"`
	Img      string         `json:"img"`
	Name     string         `json:"name"`
	Time     timeutils.Time `json:"time"`
	TimeTo   timeutils.Time `json:"time_to"`
	Duration string         `json:"duration"`
	Descr    string         `json:"descr"`
}

// Category : channel category
type Category struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}
