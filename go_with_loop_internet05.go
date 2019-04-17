package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

var washPostXML = []byte(`
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9">
<url>
<loc>
https://www.washingtonpost.com/business/technology/real-or-artificial-tech-titans-declare-ai-ethics-concerns/2019/04/07/249a1b74-5959-11e9-98d4-844088d135f2_story.html
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-07T18:36:20Z</news:publication_date>
<news:title>
Real or artificial? Tech titans declare AI ethics concerns
</news:title>
<news:keywords>
APFN-US-Safeguarding Artificial Intelligence, Business, General news, Technology, Artificial intelligence, Computing and information technology, Skin care, Personal care, Beauty and fashion, Lifestyle
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/technology/2019/04/06/they-were-settling-into-their-airbnb-then-they-found-hidden-camera/
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-06T19:48:20.076Z</news:publication_date>
<news:title>
They were settling into their Airbnb. Then they found a hidden camera.
</news:title>
<news:keywords>
airbnb, ireland, hidden camera, new zealand family, barker family, host, surveillance camera
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/business/technology/tech-entrepreneur-gives-25m-in-cryptocurrency-to-alma-mater/2019/04/05/c556bfde-57f8-11e9-aa83-504f086bf5d6_story.html
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-05T23:15:59Z</news:publication_date>
<news:title>
Tech entrepreneur gives $25M in cryptocurrency to alma mater
</news:title>
<news:keywords>
US-Cryptocurrency-University-Donation, California, United States, North America, San Francisco, Business, General news, Technology, Higher education, Education, Social affairs, Small business financing, Small business
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/business/technology/boeing-cutting-production-rate-of-troubled-737-max-jet/2019/04/05/584b7772-57e4-11e9-aa83-504f086bf5d6_story.html
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-05T23:03:28Z</news:publication_date>
<news:title>
Boeing cutting production rate of troubled 737 Max jet
</news:title>
<news:keywords>
APFN-US-Boeing-Planes-Software Fix, Business, General news, Air travel disruptions, Transportation, Aircraft manufacturing, Aerospace and defense industry, Industrial products and services, Technology, Software, Computing and information technology
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/business/technology/boeing-dealing-with-second-software-problem-on-troubled-jet/2019/04/05/47aa865c-57ca-11e9-aa83-504f086bf5d6_story.html
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-05T21:05:04Z</news:publication_date>
<news:title>
Boeing cutting production rate of troubled 737 Max jet
</news:title>
<news:keywords>
US-Boeing-Planes-Software-Fix, Business, General news, Aircraft manufacturing, Aerospace and defense industry, Industrial products and services, Technology, Aviation accidents and incidents, Transportation accidents, Accidents, Accidents and disasters, Transportation, Technology issues, Software, Computing and information technology
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/business/technology/the-latest-boeing-to-cut-production-rate-of-737-max-jet/2019/04/05/94ffeb22-57e3-11e9-aa83-504f086bf5d6_story.html
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-05T20:44:19Z</news:publication_date>
<news:title>
The Latest: Boeing to cut production rate of 737 Max jet
</news:title>
<news:keywords>
APFN-US-Boeing-Planes-Software Fix-The Latest, US Boeing Planes Software Fix, Business, General news, Technology, Air travel disruptions, Transportation, Software, Computing and information technology, Aircraft manufacturing, Aerospace and defense industry, Industrial products and services
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/technology/2019/04/05/trumps-moonshot-next-giant-leap-or-another-empty-promise/
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-05T16:24:39.540Z</news:publication_date>
<news:title>
Trump’s moonshot: The next giant leap or another empty promise?
</news:title>
<news:keywords>
nasa, moon, return to the moon, pence, white house, homer hickam, lunar landing, sls, boeing
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
<url>
<loc>
https://www.washingtonpost.com/technology/2019/04/04/die-robocalls-die-how-to-guide-stop-spammers-exact-revenge/
</loc>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2019-04-04T19:26:43.731Z</news:publication_date>
<news:title>
Die, robocalls, die: A how-to guide to stop spammers and exact revenge
</news:title>
<news:keywords>
robocalls, robocallers, robocalling, spam, nuisance, fraud, apps, spammers, STIR/SHAKEN, AT&T, T-Mobile, Verizon, Call Protect, Do Not Call, FTC, caller-ID, Hiya, Nomorobo, RoboKiller, TrueCaller, YouMail, TNS, robo-calls, spam calls, Call Filter, Scam ID, block, robocall
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
</urlset>
`)


func main() {
	var s Sitemapindex
	var n News
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	bytes := washPostXML
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}

	for idx, _ := range n.Keywords {
		news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
	}
	
	for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}
}
