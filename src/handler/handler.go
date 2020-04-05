package handler

import (
	"log"
	"os"

	urlFetcher "github.com/RyuseiNomi/delay_reporter_lm/src/fetcher"
	jsonWorker "github.com/RyuseiNomi/delay_reporter_lm/src/json"
	s3Uploader "github.com/RyuseiNomi/delay_reporter_lm/src/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	bucketName = "delay-list"
	key        = "delay-list.json"
)

func Handler() {
	delayList := urlFetcher.GetDelayList()

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
