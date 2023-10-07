package main

import (
	"math/rand"
	"time"
)

func main() {
	var listem []int
	rand.Seed(time.Now().Unix())

	for {
		go func() {
			for {
				listem = append(listem, rand.Intn(999999999)*rand.Intn(999999999))
			}
		}()

	}

}
