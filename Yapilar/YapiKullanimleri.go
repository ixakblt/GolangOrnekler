/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/
// Yapıların kullanım şekilleri
package main

import "fmt"

func main() {
	a := ana{
		test1: test1{
			isim: "a",
		},
	}
	fmt.Printf("%v\n", a)
	b := ana{}
	b.anastc = "ss"
	b.isim = "keklik"
	fmt.Printf("%v\n", b)

	c := ana{}
	c.test1.isim = "hi all"
	fmt.Printf("%v\n", c)

	d := ana{}
	d.test1 = test1{
		isim: "denemeee",
	}
	fmt.Printf("%v\n", d)
	//
	//
}

type test1 struct {
	isim string
}

type ana struct {
	anastc string
	test1
}
