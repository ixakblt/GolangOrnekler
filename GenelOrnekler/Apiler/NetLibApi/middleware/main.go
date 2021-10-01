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
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()              // Mux nesnesi ürettik
	mux.HandleFunc("/", AradakiAdam(Sido)) // dizini dinlemeye aldık
	http.ListenAndServe(":8080", mux)      // serveri ayağa kaldırdk ve kendi mux nesnemizi veri olarka yolladık
}

func Sido(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html") // htmp dosyasını parçaladık
	a := make(map[string]interface{})            // göndereceğimiz veriler için map yaptık
	a["isim"] = "selam knk "                     // map e ekleme yaptık
	tmpl.Execute(w, a)                           //map imizi index e gönderdik index içerisinde {{"isim"}} formatı ile karşıladık bu html template özelliği
}

//aradaki adamı yazdık middleware bu arkadaş
func AradakiAdam(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Time : %v istek geldi", time.Now()) // istediğimiz işlemi bu araya yazacaz ben konsola print basmayı tercih ettim basit olması adına
		next(w, r)                                      //logumuzu tutuk şimdi istği geri yollayaılım
	}
}

/*
middleware ile gelen istekleri filtreleyip işlemler yapılabilir içerisinde if yapıları kullanılarak yönlendirmeler yapılanilir.
*/
