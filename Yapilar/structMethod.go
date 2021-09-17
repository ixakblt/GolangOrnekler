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

import "fmt" // standart çıktı için gereken paket

// veriler adında in tipinde 2 değer alan bir struct tanımladık
type veriler struct {
	veribir int
	veriki  int
}

// veriler adında ki struct için method yazdık gelen verileri toplyarak return ediyor.
func (v veriler) topla() int {
	return v.veribir + v.veriki
}

func main() {
	// tanımlama yaptık
	girdiler := veriler{}
	//değerleri atadık
	girdiler.veribir = 11
	girdiler.veriki = 55
	//girdiler adında tanımladığımız veriler tipinde ki struct için yazdığımız methodu ekrana yazdırdık
	fmt.Println(girdiler.topla())
}
