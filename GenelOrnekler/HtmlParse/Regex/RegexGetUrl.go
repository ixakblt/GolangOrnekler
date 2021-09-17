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
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, _ := http.Get("https://fb.com")   //request attık
	alis, _ := ioutil.ReadAll(resp.Body)    //veriyi okuduk
	re := regexp.MustCompile(`http([^"]*)`) //düzenli ifademizi oluşturduk
	veri := re.FindAll(alis, -1)            //ifade ile aramamızı yaptık 2. değer kaç döndü istediğimizdir biz hepsini almak için -1 dedik
	for _, a := range veri {                // döngü ile ekrana bastırdık
		fmt.Println(string(a))
	}

}
