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

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()                      //Fiber nesnesi tanımladık
	app.Get("/", func(c *fiber.Ctx) error { //Gelen get isteğini iç içe fonksiyon ile karşıladık
		return c.SendString("selam knk ilk api") // String gönderdik
	})

	app.Static("/", "./") //alan.adi/ itibari ile bulunduğumuz
	//klasörü erişime açtık bu site yazarken js css gibi dosyalara erişim için gerekli

	app.Get("Down", harici) //Parantez kullanmadan dışarda yazıdığımız fonksiyonu çağırdık proje büyüdükçe routerlar ayrılacak orda bu method kullanılır.

	app.Listen(":8080") //8080 portunu dinlemeye aldık api ayağa kalktı
}

func harici(c *fiber.Ctx) error {
	return c.Download("go.mod")
}
