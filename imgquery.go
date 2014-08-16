package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type PageSetting struct {
	BaseUrl    string       `json:"baseUrl"`
	Query      string       `json:"query"`
	FindKey    string       `json:"findKey"`
	ImgFindKey string       `json:"imgFindKey"`
}

func (p PageSetting) Url(keyword string) string {
	return p.BaseUrl + p.Query + keyword
}

func (p PageSetting) GetImagePaths(keyword string) ([]string) {
	url := p.Url(keyword)
	fmt.Printf("startUrl = %s\n", url)
	if p.FindKey != "" {
		return newDocumentUrl(url, p.FindKey, p.ImgFindKey)
	} else {
		return newDocumentImage(url, p.ImgFindKey)
	}
}

func newDocumentUrl(url, findKey, imgFindKey string) []string {
	var doc *goquery.Document
	var err error
	if doc, err = goquery.NewDocument(url); err != nil {
		return []string{}
	}

	result := make([]string, 0)
	doc.Find(findKey).Each(func(_ int, s *goquery.Selection) {
		path, exists := s.Attr("href")
		if !exists {
			return
		}
		fmt.Printf("url = %s\n", path)
		result = append(result, newDocumentImage(path, imgFindKey)...)
	})
	return result
}

func newDocumentImage(url, findKey string) []string {
	var doc *goquery.Document
	var err error
	if doc, err = goquery.NewDocument(url); err != nil {
		return []string{}
	}

	result := make([]string, 0)
	doc.Find(findKey).Each(func(_ int, s *goquery.Selection) {
		path, exists := s.Attr("src")
		if exists {
			result = append(result, path)
		}
	})
	return result
}
