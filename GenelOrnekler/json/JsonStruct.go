/*
written by raifpy

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type apiResType struct {
	Data struct {
		Location  string `json:"location"`
		Confirmed int    `json:"confirmed"`
		Deaths    int    `json:"deaths"`
		Recovered int    `json:"recovered"`
		Active    int    `json:"active"`
	} `json:"data"`
	Dt string  `json:"dt"`
	Ts float64 `json:"ts"`
}

var api = "https://covid2019-api.herokuapp.com/v2/country/" + os.Args[1] // turkey

func main() {
	response, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var veri apiResType
	ham, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(ham, &veri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Country:  ", veri.Data.Location)
	fmt.Println("Deaths:   ", veri.Data.Deaths)
	fmt.Println("Confirms: ", veri.Data.Confirmed)

}
