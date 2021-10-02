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
	a := make(chan bool, 1) // chan oluşturduk bool taşıyor
	go func(msg string) {   // msg değişkeni alacak asenk vi fon kurduk

		fmt.Println(msg) // bla bla işem yaptık
		a <- true        // işlem bitti true döndük chan verimize
	}("going")

	<-a //true gelince program kapndı
}
