package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	//fmt.Println("Go Colly is run")
	c := colly.NewCollector(colly.AllowedDomains("www.jackjones.com"))
	c.Visit("https://www.jackjones.com/nl/en/jj/shoes/")
	c.OnRequest(func(r *colly.Request) { fmt.Println("Scrapping:", r.URL) })
	c.OnResponse(func(r *colly.Response) { fmt.Println("Scrapping:", r.StatusCode) })
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})
}
