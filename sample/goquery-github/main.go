package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// イベントページ取得
	doc, err := goquery.NewDocument("https://github.com/trending")
	if err != nil {
		return
	}

	// リポジトリ部分抜き出し
	//  root: <ol class="repo-list">
	olSelection := doc.Find("ol.repo-list > li")
	olSelection.Each(func(_ int, s *goquery.Selection) {
		// log.Println(format(nil, s.Find("div.f6.text-gray.mt-2 > a").Text()))
		// log.Println("\n--------------------------")
		// log.Println(format(nil, s.Find("div.f6.text-gray.mt-2 > a").First().Text()))
		// log.Println("\n--------------------------")
		log.Println("[detail]--------------------------")
		s.Find("div.f6.text-gray.mt-2 > a").Each(func(_ int, s2 *goquery.Selection) {
			//log.Println(s2.Text())
			href, _ := s2.Attr("href")
			log.Println("href: " + href)
			if strings.Contains(href, "stargazers") {
				// スターの部分取得
				log.Println(format(nil, s2.Text()))
			}
		})
		log.Println("[detail]--------------------------")
	})
}

// データクレンジングメソッド
func format(r *regexp.Regexp, target string) string {
	result := strings.Replace(target, "\n", "", -1)
	result = strings.Replace(result, " ", "", -1)
	// result = strings.Replace(result, "定員", "", -1)
	// result = r.ReplaceAllString(result, "")
	// result = strings.Replace(result, "／", "/", -1)

	return result
}
