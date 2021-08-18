package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

// Site to crawl
const site = "https://twitter.com/david_bujosa/status/1165109248868724737"

type tweet struct {
	Name     string
	Username string
	Message  string
}

// Note: This is example is outdated because the Twitter frontend has changed
func main() {
	c := colly.NewCollector()

	messages := []tweet{}

	c.OnHTML(".tweet", func(e *colly.HTMLElement) {
		messages = append(messages, tweet{
			Name:     e.ChildText(".account-group .fullname"),
			Username: e.ChildText(".account-group .username"),
			Message:  e.ChildText(".tweet-text"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(messages))
}
