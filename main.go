package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
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

// 構造体を定義
type DataItem struct {
	Question string
	ID       string
	Answer   string // 回答を格納するフィールドを追加
}

func main() {
	url := "https://docs.google.com/forms/d/e/1FAIpQLSfoEtgxNENKW3cK9Nfm-z7RGtEcUbdrrKxuyZkOveDUykgR-w/viewform"

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

		qRegex := regexp.MustCompile(`&#34;(.*?)&#34;`)
		qMatch := qRegex.FindString(match[1])
		if qMatch != "" {
			qList = append(qList, strings.ReplaceAll(qMatch, "&#34;", ""))
		}
	}

	fmt.Println(idList)
	fmt.Println(qList)

	// 以下、おかしすぎる
	dataList := make([]DataItem, len(qList))
	reader := bufio.NewReader(os.Stdin)
	for i := range qList {
		dataList[i] = DataItem{Question: qList[i], ID: idList[i]}
		fmt.Printf("'%s' : ", qList[i])
		answer, _ := reader.ReadString('\n')
		dataList[i].Answer = strings.TrimSpace(answer)
	}

	var payload [][]string

	for _, item := range dataList {
		payload = append(payload, []string{item.ID, item.Answer})
	}

	fmt.Println(payload)

	// これをPOSTしてしまえばいいのです。
}
