package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func scraperCMD() error {
	url := "https://www.larvalabs.com/cryptopunks/details/7804"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".header_block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Find("p").Each(func(ii int, ss *goquery.Selection) {
			fmt.Println(ss.Text())
		})
	})
}
