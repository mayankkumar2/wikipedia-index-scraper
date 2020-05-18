package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"wikipediaScraper/pkgs/utils"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	index := utils.MakePermutations()
	f, err := os.Create("./data")
	check(err)
	defer f.Close()
	c := colly.NewCollector()
	c.OnHTML("a.mw-redirect",  func(e *colly.HTMLElement) {
		f.WriteString(fmt.Sprintf("https://en.wikipedia.org%s\n",e.Attr("href")))
	})
	for _, v := range index {
		c.Visit(fmt.Sprintf("https://en.wikipedia.org/wiki/Special:AllPages/%s",v))
		fmt.Println("Wrote file for : ",v)
	}
}
