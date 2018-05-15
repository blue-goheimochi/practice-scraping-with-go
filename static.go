package main

import (
	"fmt"
	"strings"
	
	"github.com/PuerkitoBio/goquery"
)

func main () {
	url := "https://www.nikkei.com/"

	doc, _ := goquery.NewDocument(url)
	fmt.Println("--- <title>タグ ----------")
	fmt.Println(doc.Find("title").Text())
	fmt.Println("");

	fmt.Println("--- アクセスランキング ----------")
	doc.Find("ul.m-sub_access_ranking > li").Each(func(i int, s *goquery.Selection) {
		rank := s.Find("a > span:nth-child(1)").Text()
		title := s.Find("a > span.m-sub_access_ranking_title").Text()
		fmt.Println(strings.TrimSpace(rank) + " " + fixTitle(title))
	})
}

func fixTitle (title string) string {
	return strings.TrimSpace(title)
}