package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
[5 5]type == array
[]type == slice
*/

// take data inside CD
type cdIndex struct {
	CdInfos []CdInfo `xml:"CD"`
}

// Read data inside xml
type CdInfo struct {
	Title  string `xml:"TITLE"`
	Artist string `xml:"ARTIST"`
	Price  string `xml:"PRICE"`
}

// convert struct to string
func (l CdInfo) String() string {
	return fmt.Sprintln(l.Title, l.Price, l.Artist)
}

func main() {
	resp, _ := http.Get("https://www.w3schools.com/xml/cd_catalog.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var s cdIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.CdInfos)
	for _, CdInfo := range s.CdInfos {
		fmt.Println("Prince of :", CdInfo.Title, ", written by :", CdInfo.Artist, ", is: ", CdInfo.Price)
	}
}

/*
<CD>
<TITLE>Empire Burlesque</TITLE>
<ARTIST>Bob Dylan</ARTIST>
<COUNTRY>USA</COUNTRY>
<COMPANY>Columbia</COMPANY>
<PRICE>10.90</PRICE>
<YEAR>1985</YEAR>
</CD>
<CD>
<TITLE>Hide your heart</TITLE>
<ARTIST>Bonnie Tyler</ARTIST>
<COUNTRY>UK</COUNTRY>
<COMPANY>CBS Records</COMPANY>
<PRICE>9.90</PRICE>
<YEAR>1988</YEAR>
</CD>
*/
