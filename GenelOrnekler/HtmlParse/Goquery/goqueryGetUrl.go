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
	var linkler []string //string listesi oluşturduk
	siteurl := "https://fb.com"
	resp, _ := http.Get(siteurl)                           // get isteği attık
	doc, _ := goquery.NewDocumentFromReader(resp.Body)     //goquery nesnesi oluşturduk
	doc.Find("a").Each(func(i int, s *goquery.Selection) { // tüm a ları işleme aldık
		basliklar, _ := s.Attr("href")       // href değerlerini aldık
		linkler = append(linkler, basliklar) //listemize ekledik

	})
	for i := range linkler { // for ile liste uzunluğunda döngü kurduk burda index değerini tek aldık
		if len(linkler[i]) < 1 || linkler[i] == "#" { // link olup olmadığını kontrol ettik
			fmt.Printf("%v link degil\n", linkler[i])
		} else {
			if len(linkler[i]) > 2 && linkler[i][:2] != "ht" { // eğer ht ile başlamıyorsa başına url ekledik
				if linkler[i][0] != '/' {
					veri := siteurl + "/" + linkler[i]
					fmt.Println(veri)
				} else {
					veri := siteurl + linkler[i]
					fmt.Println(veri)
				}

			} else {
				fmt.Println(linkler[i])
			}

		}

	}
}
