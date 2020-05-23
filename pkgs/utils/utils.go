package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"regexp"
	"strconv"
	"wikipediaScraper/pkgs/sets"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadMap(m *map[string] []string){
	f, _ := os.Open("graph")
	st := make([]byte, 9012400)
	f.Read(st)
	json.Unmarshal(st, m)
}
func charRange(min uint8, max uint8) []uint8 {
	if max < min {
		return []uint8{}
	}
	s := make([]uint8, max-min)
	for i := 0 ; min < max ; i++{
		s[i] = min
		min++
	}
	return s
}

func MakePermutations() []string {
	upperCaseCh := charRange('A','Z'+1)
	lowerCaseCh := charRange('a','z'+1)
	CharSet := make([]string,0, 26*26*2)
	for _, v := range upperCaseCh {
		for  _,v2 := range upperCaseCh {
			CharSet = append(CharSet,(string([]byte{byte(v), byte(v2)})))
		}
	}
	for _, v := range upperCaseCh {
		for  _,v2 := range lowerCaseCh {
			CharSet = append(CharSet,(string([]byte{byte(v), byte(v2)})))
		}
	}
	return CharSet
}


func scr() {
	index := MakePermutations()
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


func IndexOf(element string, data []string) (int) {
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
	s := make(map[string][]string)
	var se *sets.Set
	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		m, err := regexp.MatchString("^/wiki/.*", e.Attr("href"))
		check(err)

		if m {
			val := IndexOf("https://en.wikipedia.org"+e.Attr("href"), d)
			if val != -1 {
				se.Add(strconv.Itoa(val))
			}
		}
	})
	for i, v := range d {
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