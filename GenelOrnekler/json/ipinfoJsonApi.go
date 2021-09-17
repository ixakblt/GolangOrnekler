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
	"io/ioutil"
	"net/http"
)


var elitveri map[string]interface{}
func main() {
	var ip string
	fmt.Printf("ip veya domain giriniz (https olmadan): ")
	fmt.Scan(&ip)
	urlimiz := fmt.Sprintf("http://api.ipapi.com/%v?access_key=f76460f4e8673c445e56d09769f54243&output=json",ip)
	veri ,_:= http.Get(urlimiz)
	a ,_:=ioutil.ReadAll(veri.Body)
	json.Unmarshal(a ,&elitveri )
	fmt.Printf("\n\nIp Adresi : %v\nKita : %v\nUlke : %v\nSehir : %v\nKonum : %v , %v \n\n--ElitHatz--",elitveri["ip"],elitveri["continent_name"],elitveri["country_name"],elitveri["city"],elitveri["latitude"],elitveri["longitude"])
}
