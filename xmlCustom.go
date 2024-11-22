package excelize

import "encoding/xml"

type CustomProperties struct {
	Pairs map[string]string
}

type decodeCustomProperties struct {
	XMLName    xml.Name          `xml:"http://schemas.openxmlformats.org/officeDocument/2006/custom-properties Properties"`
	Vt         string            `xml:"xmlns vt,attr"`
	Properties []*customProperty `xml:"property"`
}

type customProperty struct {
	Fmtid    string `xml:"fmtid,attr"`
	Pid      string `xml:"pid,attr"`
	Name     string `xml:"name,attr"`
	Lpwstr   string `xml:"lpwstr"`
	Bool     string `xml:"bool"`
	Filetime string `xml:"filetime"`
	R8       string `xml:"r8"`
}
