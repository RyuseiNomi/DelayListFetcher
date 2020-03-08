package handler

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	url = "https://tetsudo.rti-giken.jp/free/delay.json"
)

func Handler() {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	log.Printf(string(byteArray))
}
