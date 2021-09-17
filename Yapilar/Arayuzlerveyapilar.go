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

type veri struct {
	Aa int
} //yapı tanımladık

// Yap fonk tanımladık
func (v veri) Yap() {
	fmt.Println("selam")
}

//Selam interfeys yazdık yap fonk u olması gerekli dedil
type selam interface {
	Yap()
	Setet()
	Getmsg() int
}

func (v *veri) Setet() {
	// adresine set ettik eğer pointer kullanmazsak var olan nesne etkilenmez yeni nesne oluşturulur bunu da return etmek gerekir
	//veri return edecek bir fonksiyon olmalıydı pointer olmasaydı ve son hali return edilip değişken ile karşılanmalıydı
	v.Aa = 4545

}

//Get Fonk yazdık veriyi return ediyor
func (v *veri) Getmsg() int {
	return v.Aa
}

func main() {
	var kk selam       // interfeys tanımladık
	ss := veri{Aa: 55} // yapıdan değişken türettik
	kk = &ss           // yapıyı interfeys e bağladık
	fmt.Println(kk)    // çıktımıza baktık kk içinde 55 var
	kk.Yap()           //ss için yazılmış methodu kk için kullandık
	// kk.Aa  = şeklinde bir kullanım yok interfeys üzerinden yapıya veri yazamayız bunun için set fonk u yazılır
	kk.Setet()      // Set ettik
	fmt.Println(kk) // kontrol ettik
	fmt.Printf("%v \n", ss.Aa)
	ps := kk.Getmsg() //getmsg ile değişkene aldık veriyi
	fmt.Println(ps)   //kontrol ettik

}
