/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
Yapılar ,yapıcı method, basitçe strings paketi
 return değeri formatlama ve yapı için method
 işlemini kısaca basit bir örnek ile açıklamaya çalıştım
*/

package main //paket adı klasik olarak ☺

import (
	"fmt"     //olmazsa olmazımız fmt paketi
	"strings" // Harfleri büyütmek için gereken strings paketi
)

type veri struct { //yapı tanımladık
	deger1, deger2 int // değerleri ekledik
}

//Yapıcı fonksiyonu yazdık
func olustur() *veri { // Pointer olarak dönüş yapmakta
	a := new(veri) // new ile nesne oluşturup a değişkenine atadık
	return a
}

//Metod tanımladıık
func (v veri) Topla(a string) string {
	return fmt.Sprintf("Sayın %s Hanım Değerlerin Toplamı :  %d ' dir. ", strings.ToUpper(a), (v.deger1 + v.deger2))
	/*Asıl işimiz burda ☺ a olarak aldığımız string veriyi ToUpper fonksiyonunu da kullanarak %s ile string içine koyduk
	Ardından değerlerin toplamı için de %d ile decimal bir formatlama kullandık
	Parantez içinde değerleri toplayıp a ve toplamı string içine gönderdik
	direkt return etmek istediğimizde kullandığımız virgül karakteri gonun
	birden fazla return özelliğinden dolayı sorun çıkartıyor bruda imdadımıza
	fmt içinde ki Sprintf kardeşimiz yetişiyor ☺ printf de olduğu gibi burda da f var
	yani formatlama için kullanılacağını anlıyoruz
	*/

}

//main func
func main() {
	islem := olustur() //yukarda tanımladğımız yapıcı metod ile islem mesnesini türettik
	islem.deger1 = 23  //değer ataması falan ☺
	islem.deger2 = 24
	fmt.Println(islem.Topla("Trinity"))
	/*nesnemizin Topla metodunu çağırdık içerde ki a string değişkenine de değerimizi verdik
	  ve çalıştırdığımızda çıktımız  "Sayın TRINITY Hanım Değerlerin Toplamı :  47 ' dir. " ieklinde oalcak
	*/
}
