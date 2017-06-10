package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"strings"
)

func NewPunk(ID int) (*Punk) {
	url := fmt.Sprintf("https://www.larvalabs.com/cryptopunks/details/%d", ID)
	fmt.Println(url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return &Punk{
		ID: ID,
		Type: Type(doc),
		Owned: isOwned(doc),
		Accessories: Accessories(doc),
	}
}

func Type(doc *goquery.Document) (rightType string) {
	doc.Find(".header_block").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(ii int, ss *goquery.Selection) {
			if strings.Contains(ss.Text(), "other punks like this") {
				xpd := strings.Split(strings.TrimSpace(ss.Text()), "-")
				rightType = strings.TrimSpace(xpd[0])
			}
		})
	})
	return
}

func isOwned(doc *goquery.Document) (owned bool) {
	owned = false
	doc.Find(".header_block").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(ii int, ss *goquery.Selection) {
			if strings.Contains(ss.Text(), "This punk is currently owned by address") {
				owned = true
			}
		})
	})
	return
}

func Accessories(doc *goquery.Document) (accessories []string) {
	doc.Find(".header_block").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(ii int, ss *goquery.Selection) {
			xpd := strings.Split(strings.TrimSpace(ss.Text()), "-")
			accessories = append(accessories, strings.TrimSpace(xpd[0]))
		})
	})
	return
}
