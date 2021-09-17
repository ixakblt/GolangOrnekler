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
	"io"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func upload(c echo.Context) error {

	bot, _ := tgbotapi.NewBotAPI("1498255154:AAH0wKk_kasdiqyiphQ_cjiCVTz0nasdsdHZ24") //bot token
	bot.Debug = true
	var ChatID int64 = 1316224975 // çet aydi
	var dosyalar []string         // gönderilen dosyalar
	//formdan gelen çoklu nesneleri aldık
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"] // files name değerine sahip olanlari aldık

	for _, file := range files { // döngüde gezip her biri için işlem yaptık
		src, err := file.Open() //dosyayı açıp src ye atadık
		if err != nil {
			return err
		}
		defer src.Close() //iş bitince kapat dosyayı dedik

		// Destination
		dst, err := os.Create(file.Filename) //dosya adı ile aynı simde dosya oluşturduk
		if err != nil {
			return err
		}

		dosyalar = append(dosyalar, file.Filename) //dosya adını listemize ekledik

		if _, err = io.Copy(dst, src); err != nil { //src içeriğini dst ye yazdık
			return err
		}

		msg := tgbotapi.NewDocumentUpload(ChatID, file.Filename) //dosya adını verdiğimiz id ye yollamak için değişkeni ayarladık
		bot.Send(msg)                                            // mesaj piyuwww
		dst.Close()                                              // dosyayı kapat
	}

	for _, key := range dosyalar { //dosyalar gitti sıra silmede listeden adlarını okuyup sildik

		os.Remove(key)

	}
	c.Redirect(http.StatusPermanentRedirect, "https://t.me/ixhayabusa") // iş bitti şimdigit telegram adresime yönlendir baba heyran
	return nil
}

func main() {
	e := echo.New() // nesne türettik

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")   // heştiyemel klasörümüzü belirledik
	e.POST("/upload", upload) //api endpoint belirledik

	e.Logger.Fatal(e.Start(":1323")) //verdiğimiz portta serveri kaldırdık
}
