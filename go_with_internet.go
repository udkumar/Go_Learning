package main

// we can use the encoding/xml package to parse xml
// To print the data out, we'll use fmt,
// to make the request we'll use net/http and
// to read the data, we'll need to import io/ioutil
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// we're going to build a slice of URLs. We'll then make a struct for our data like
type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String () string {
  return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	// we want to access the body within, which we can find by resp.Body. 
	// That said, we will need to read this in with ioutil.ReadAll():
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	var s Sitemapindex

	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	// we can just simply convert the body data to a string with string()
	//string_body := string(bytes)
	//fmt.Println(string_body)

	// Don't forget to close to free up resources:
       //	resp.Body.Close()
}
