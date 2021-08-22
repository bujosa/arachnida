package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

// Site to crawl
const site = "https://en.wikipedia.org/wiki/Bad_Bunny"

type wikipedia struct {
	Title      string
	Picture    string
	Born       string
	Occupation string
	Genres     string
}

func main() {
	c := colly.NewCollector()

    page := wikipedia{}

	c.OnHTML(".infobox", func(e *colly.HTMLElement) {
		
		page.Title = e.ChildText("#mw-content-text > div.mw-parser-output > table.infobox.biography.vcard > tbody > tr:nth-child(1) > th > div")
		page.Picture = e.ChildAttr("#mw-content-text > div.mw-parser-output > table.infobox.biography.vcard > tbody > tr:nth-child(2) > td > a > img", "src")
		page.Born = e.ChildText("#mw-content-text > div.mw-parser-output > table.infobox.biography.vcard > tbody > tr:nth-child(3) > td")
		page.Occupation = e.ChildText("#mw-content-text > div.mw-parser-output > table.infobox.biography.vcard > tbody > tr:nth-child(4) > td")
		page.Genres = e.ChildText("#mw-content-text > div.mw-parser-output > table.infobox.biography.vcard > tbody > tr:nth-child(8) > td > div > ul")

	})

	err := c.Visit(site)

	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(page, "", "\t")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}