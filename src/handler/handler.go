package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

	if err := createJSON(delayList); err != nil {
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

/**
 * Create empty JSON file on temp file and write fetched delay-list
 */
func createJSON(delayList []byte) error {
	if delayList == nil {
		return fmt.Errorf("create JSON error: %s", "nil bytes was given")
	}

	if err := os.MkdirAll(tempDir, 0777); err != nil {
		return err
	}

	file, err := os.Create(tempDir + key)
	if err != nil {
		return err
	}

	_, err = file.Write(delayList)
	if err != nil {
		return err
	}

	if isExist := isExistTempFile(tempDir); isExist != true {
		return fmt.Errorf("Temp file does not exist")
	}

	return nil
}

func isExistTempFile(tempFile string) bool {
	_, err := os.Stat(tempFile)
	return !os.IsNotExist(err)
}
