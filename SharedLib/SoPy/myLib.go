package main //paket adi main olmak zorunda

import "C" //cgo olarak kullanmamıza ayrayan kütüphane

func main() {}

//export YazDostum
func YazDostum() *C.char {
	a := "Hi Github i am ixakblt this test So file"
	return C.CString(a)
}

/*
Go dilinde Fonksiyon açıklamaları önemlidir burda ki export YazDostum açıklaması export etmemize olanak sağlıyor
go build -buildmode=c-shared -o myLib.so myLib.go
komutu ile derlememizi gerçekleştiriyoruz.

*/
