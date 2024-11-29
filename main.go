package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	return doc.Html()
}

func main() {
	url := "url"

	sendURL := regexp.MustCompile(`viewfor.+`).ReplaceAllString(url, "formResponse")
	fmt.Println("url : " + sendURL)

	siteHTML, err := fetch(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	var idList []string
	var qList []string

	// すべて取得
	re := regexp.MustCompile(`%\.@\.\[(\d.*?)\]`)
	allList := re.FindAllStringSubmatch(siteHTML, -1)

	for _, match := range allList {
		fmt.Println(match[1])
	}

	fmt.Println("================")

	// 取得
	for _, match := range allList {
		idRegex := regexp.MustCompile(`^\d+`)
		id := idRegex.FindString(match[1])
		if id != "" {
			idList = append(idList, id)
		}

		qRegex := regexp.MustCompile(`\d+,"(.*?)"`)
		qMatch := qRegex.FindStringSubmatch(match[1])
		if len(qMatch) > 1 {
			qList = append(qList, strings.ReplaceAll(qMatch[1], "\"", ""))
		}
	}

	fmt.Println(qList)
	fmt.Println(idList)
}
