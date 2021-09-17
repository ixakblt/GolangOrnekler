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

	"github.com/labstack/echo"
)

func main() {
	// nesne üretik
	e := echo.New()
	//ana dizinine gelen get isteiğini karşıladık
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "say may neym")
	})
	// adimne path ine gelen istekelrden kedileri sevin sorgusunu karşıladık gelen verii formatlayarak yeniden kullandık
	e.GET("/adimne", func(c echo.Context) error {
		benimadimarda := c.QueryParam("kedilerisevin")
		retet := fmt.Sprintf("emrin %v yerine getirildi", benimadimarda)
		return c.String(200, retet)

		//http://localhost:3200/adimne?kedilerisevin=miyav
	})
	e.GET("/jesonver", jesonver) //farklılık olsun fonksiyonu ayrı yerde yazıp buraya veriyi çekelim

	//ayağa kaldırdık
	e.Logger.Fatal(e.Start(":3200"))

}

//Ver bana json fonksiyonu
func jesonver(c echo.Context) error {
	iboo := &Yaknknaber{
		Isim:    c.QueryParam("isim"), //Sorguda gelen veriler tek tek karşılayıp yapının içine tanımladık
		Soyiism: c.QueryParam("patronu"),
	}
	return c.JSON(http.StatusOK, iboo) //json olması için ayarladığımız yapıyı return ettik
	//http://localhost:3200/jesonver?isim=Neo&patronu=Trinity
}

//dönecek json yapısını tanımladık string yazısından sonrası json çıktısında değerlerin keyleri olacak
type Yaknknaber struct {
	Isim    string `json:"isim"`
	Soyiism string `json:"Patron"`
}
