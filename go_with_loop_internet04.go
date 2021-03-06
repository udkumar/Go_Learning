package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	 resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	 bytes, _ := ioutil.ReadAll(resp.Body)
	 var s Sitemapindex
	 xml.Unmarshal(bytes, &s)
	 for _, Location := range s.Locations {
	 	fmt.Printf("%s\n", Location)
	 }
	//fmt.Println(s.Locations)

}
