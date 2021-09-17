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

	yer := make(map[string]string) // Key ve Val string olan bir map oluşturduk
	yer["isim"] = "Alice"          // isim keyine aliyi atadık
	fmt.Println(yer["isim"])       // isim sorgusunun çıktısını istedik çıktı olarak Alice değerini aldık

}
