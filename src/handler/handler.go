package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	url        = "https://tetsudo.rti-giken.jp/free/delay.json"
	tempDir    = "/temp/json/delay-list.json"
	bucketName = "delay-list"
	key        = "delay-list.json"
)

func Handler() {
	delayList := getDelayList()

	if err := createJSON(delayList); err != nil {
		fmt.Errorf("Create JSON Error: %s", err)
	}

	// S3接続インスタンスの生成
	credential := credentials.NewStaticCredentials("dummydummydummy", "dummydummydummy", "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: credential,
		Region:      aws.String("ap-notheast-1"),
	})
	//svc := s3.New(session)

	if err := uploadJSON(sess); err != nil {
		fmt.Errorf("Upload JSON Error: %s", err)
	}

	log.Printf("Success to upload delay list")
}

func getDelayList() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	delayList, _ := ioutil.ReadAll(resp.Body)

	return delayList
}

func deleteExistFile() {
	//TODO 既存のJSONを削除する処理の実装
}

func createJSON(delayList []byte) error {
	if delayList == nil {
		return fmt.Errorf("create JSON error: %s", "nil bytes was given")
	}

	if err := os.MkdirAll(tempDir, 0777); err != nil {
		return err
	}

	file, err := os.Create(tempDir)
	if err != nil {
		return err
	}

	file.Write(delayList)

	return nil
}

func uploadJSON(sess *session) error {
	jsonFile, err := os.Open(tempDir)
	if err != nil {
		fmt.Errorf("upload JSON error: %s", err)
	}
	defer jsonFile.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   jsonFile,
	})

	if err != nil {
		return err
	}

	return nil
}
