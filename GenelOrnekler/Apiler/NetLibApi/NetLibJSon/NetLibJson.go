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
	"encoding/json"
	"fmt"
	"net/http"
)

type isimler struct {
	Isim    string
	Soyisim string
}

func main() {

	a := isimler{Isim: "selam", Soyisim: "naber knk"}
	s, _ := json.Marshal(a)
	fmt.Println(string(s))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path[1:])
		fmt.Fprintf(w, "%v", string(s))
	})
	http.ListenAndServe(":5555", nil)

	fmt.Println("Web Sunucu")
}
