package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	jsonWorker "github.com/RyuseiNomi/delay_reporter_lm/src/json"
	s3Uploader "github.com/RyuseiNomi/delay_reporter_lm/src/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	url        = "https://tetsudo.rti-giken.jp/free/delay.json"
	tempDir    = "/tmp/json/"
	bucketName = "delay-list"
	key        = "delay-list.json"
)

func Handler() {
	delayList := getDelayList()

	if err := jsonWorker.CreateJSON(delayList); err != nil {
		log.Fatal("Create JSON Error:  %+v\n", err)
	}

	jsonFile, err := os.OpenFile("/tmp/json/delay-list.json", os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("Open JSON File Error:  %+v\n", err)
	}
	defer jsonFile.Close()

	uploader := s3Uploader.GetUploader()

	// Upload File With Custom Session
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   jsonFile,
	})
	if err != nil {
		log.Fatal("(File Upload Error:  %+v\n", err)
	}

	log.Printf("Success to upload delay list")
}

/**
 * Fetch Delay List as a JSON File from Web site
 */
func getDelayList() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	delayList, _ := ioutil.ReadAll(resp.Body)

	log.Printf("Succeeded to get Delay-list!")
	return delayList
}
