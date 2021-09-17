/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	resp, _ := http.Get("https://go.dev/blog/")

	// goquery dökümanına çevir
	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	// liste oluştur

	basliklar := ""
	doc.Find("p a").Each(func(i int, s *goquery.Selection) {
		basliklar += "- " + s.Text() + "\n"

	})
	ss := doc.Add("p a")
	doc.Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%v %v", i, s.Text())
	})
	fmt.Println(ss.Text())
	fmt.Println("--------------------")
	fmt.Printf(basliklar)
}
