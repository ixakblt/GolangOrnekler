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
	"net/http"
)

func main() {
	mux := http.NewServeMux()         // Mux nesnesi ürettik
	mux.HandleFunc("/", Sido)         // dizini dinlemeye aldık
	http.ListenAndServe(":8080", mux) // serveri ayağa kaldırdk ve kendi mux nesnemizi veri olarka yolladık
}

func Sido(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://t.me/ixhayabusa", 301) // Yönlendirme yaptık
}
