/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/

package ixpackage //paket adını tanımladık

import "fmt" // klasik io işlemleri için gerekli paket

/*
Golang içerisinde paketler klasör içerisinde yazılır paket adı ile klasör adı aynı olması gereklidir
ixpackage adında bi klasör ve içerisinde de go dosyamızı oluşturduk ve kodlarımızı yazdık
çağırma işlemi main.go dosyasında.
func main olmaığına dikkat ediiz
*/

//Seraph paket içerisinde bi fonksiyon yazdık bu bizim testi görebilmemiz için ekrana çıktı vericek
func Seraph() {
	fmt.Println("Wingless Angel Seraph")
}
