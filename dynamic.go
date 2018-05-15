package main

import (
	"fmt"
	"strings"
	
	"github.com/PuerkitoBio/goquery"
)

func main () {
	url := "https://medium.com/"
	// url := "http://localhost:8050/render.html?url=https://medium.com&timeout=10&wait=3"

	doc, _ := goquery.NewDocument(url)
	fmt.Println("--- <title>タグ ----------")
	fmt.Println(doc.Find("title").Text())
	fmt.Println("");

	fmt.Println("--- Popular on Medium ----------")
	doc.Find("aside.extremeAdaptiveSection > ol > li").Each(func(i int, s *goquery.Selection) {
		rank, _ := s.Find("div:nth-child(1)").Html()
		title := s.Find("div:nth-child(2) > a > h3").Text()
		fmt.Println(strings.TrimSpace(rank) + ". " + fixTitle(title))
	})
}

func fixTitle (title string) string {
	return strings.TrimSpace(title)
}