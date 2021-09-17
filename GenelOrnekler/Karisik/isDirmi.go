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

//kütüphane import etme
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// klasörü okuyup dizin adında bir nesneye atadık
	dizin, _ := ioutil.ReadDir(".")
	for _, a := range dizin { // dizin içerisinde range ile gezdik bu pythondaki for i in mantığına benzemekte
		if a.IsDir() { // eğer dizinse yapılacakalr
			fmt.Println(a.Name(), "dizindir")
		} else { // değilse yapılacakalr
			fmt.Printf("%s dizin degil\n", a.Name())
		}

	}
}
