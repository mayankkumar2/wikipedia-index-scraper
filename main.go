package main

import (
	"fmt"
	"wikipediaScraper/pkgs/findPath"
)
func main(){
	fmt.Println(
		findPath.Find("https://en.wikipedia.org/wiki/AA-10","https://en.wikipedia.org/wiki/BNF_(identifier)"))
}