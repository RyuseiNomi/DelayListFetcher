package urlFetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	url = "https://tetsudo.rti-giken.jp/free/delay.json"
)

/**
 * Fetch Delay List as a JSON File from Web site
 */
func GetDelayList() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	delayList, _ := ioutil.ReadAll(resp.Body)

	log.Printf("Succeeded to get Delay-list!")
	return delayList
}
