package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"wikipediaScraper/pkgs/sets"
	"wikipediaScraper/pkgs/utils"
	"os"
	"regexp"
	"bufio"
	"strconv"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scr() {
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
func GetData() []string {
	data := make([]string,340000)
	inFile, err := os.Open("data")
	if err != nil {
		fmt.Println(err.Error() + `: ` + "data")
		return []string{}
	}
	defer inFile.Close()
	i := 0
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		data[i] = (scanner.Text())
		i++
	}
	return (data)
}


func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
func main() {
	d := GetData()
	c := colly.NewCollector()
	fmt.Println(d[:1])
	s := make(map[string] []string)
	var se *sets.Set
	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		m, err := regexp.MatchString("^/wiki/.*",e.Attr("href"))
		check(err)

		if m {
			val := indexOf("https://en.wikipedia.org"+e.Attr("href"),d)
			if val != -1 {
				se.Add(strconv.Itoa(val))
			}
		}
	})
	for i,v := range d {
		se = sets.NewSet()
		fmt.Printf("[%d] Visting %s ...\n", i, v)
		c.Visit(v)
		s[strconv.Itoa(i)] = se.Slice()
		fmt.Println(se.Slice())
		break
	}
	f, err := os.Create("graph")
	check(err)
	defer f.Close()
	b, _ := json.Marshal(s)
	f.Write(b)
}
