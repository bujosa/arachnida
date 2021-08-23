# Arachnida
This is a simple project about crawling and golang using Colly.

## Packages
Offical Documentation of [Colly](http://go-colly.org/docs/examples/reddit/) 

## Example with Colly
```
func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
```

## Colly Installation
```
module github.com/x/y

go 1.14

require (
        github.com/gocolly/colly/v2 latest
)
```