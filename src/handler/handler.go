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
	delayList := getDelayList()
	log.Printf(delayList)
}

func getDelayList() string {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
