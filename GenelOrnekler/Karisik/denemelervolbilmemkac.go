package main

import (
	"fmt"
)

/*
Bu arkadaş int veya string alabilir ikisi için de çalışır
aldığı veri tipini T tipi olarak tanımlıyor ardından
T tipini aldım diyor işlem yapıyor
*/
func AA[T ~int | ~string](t T) {
	//eğer veri tipi int ise veri int yazacak değilse değil
	if fmt.Sprintf("%T", t) == "int" {
		fmt.Printf("veri %T ve değeri %v\n", t, t)
	} else {
		fmt.Printf("veri %T ve değeri \"%v\"", t, t)
	}
}

func main() {
	//fonksiyon biden çok tipi desteklediği için tipi de belirtiyoruz yukarıda kontrolü yapıyor
	AA[int](666)
	AA[string]("Köpeğimin adını kedi koydum")

}
