# wikipedia-index-scraper
The program can map out the path between 2 wikipedia pages.

###### Note: Only the pages at https://en.wikipedia.org/wiki/Wikipedia:Contents/A%E2%80%93Z_index are considered for the project. 
## Output
##### The output of the array is the order of page navigation.

- Input
```bash
	fmt.Println(
		findPath.Find("https://en.wikipedia.org/wiki/AA-10","https://en.wikipedia.org/wiki/BNF_(identifier)"))

```

```
C:\Users\Mayank-PC\go\src\wikipediaScraper>go run main.go
[https://en.wikipedia.org/wiki/AA-10 https://en.wikipedia.org/wiki/ISBN_(identifier) https://en.wiki
pedia.org/wiki/BNF_(identifier)]
```

### The project is for educational purposes
