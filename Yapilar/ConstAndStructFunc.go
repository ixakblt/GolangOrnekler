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

const You string = "Trinity"

type Kalp struct {
	hi string
}

func (k Kalp) Love(l string) {
	fmt.Printf("I Love %s\n", l)
}
func main() {
	for {
		I := Kalp{}
		I.Love(You)
	}
}
