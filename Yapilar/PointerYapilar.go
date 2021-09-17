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

import "fmt"

func main() {

	a := YapKedi("selam", "kbr") // kedi yaptık
	fmt.Println("a : ", a)
	//
	a.isimdegisduz("merhaba") //kedi içerisinde pointer olmadan güncelleme yaptık fakat sonuca işlemedi çünkü bi kopyasını alıp işlem yaptı
	fmt.Println("a : ", a)
	//
	b := a.isimdegisduz("Naber knk") //pointersiz işlemi yapabilmek için return edip yeniden atama yapmak gerekli
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	//
	//
	//
	a.isimdegispointer("ben pointer aldım") //pointer kullanıldı diye olduğu adrese işlem yapar atama olmadan değişiklilik yapar
	fmt.Println("a : ", a)
}

type kedi struct {
	sahibi string
	isim   string
}

//bana kedi ver
func YapKedi(isim, sahibi string) kedi {
	return kedi{isim: isim, sahibi: sahibi}
}

//atama yapılır fakat a kopyasına işlem yapıldığı için yukarıda etki etmez
func (k kedi) isimdegisduz(ism string) kedi {
	k.isim = ism
	return k
}

// Direkt adrese etki eder değişiklik yapar
func (k *kedi) isimdegispointer(ism string) {
	k.isim = ism
}
