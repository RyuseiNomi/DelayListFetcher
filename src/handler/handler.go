package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

/**
 * Lambda関数の入り口となる関数
 */
func Handler() {
	delayList := getDelayList()

	if err := createJSON(delayList); err != nil {
		fmt.Errorf("Create JSON Error: %s", err)
	}

	jsonFile, err := os.Open(tempDir)
	defer jsonFile.Close()
	if err != nil {
		fmt.Errorf("Upload JSON Error: %s", err)
	}

	// S3接続インスタンスの生成
	credential := credentials.NewStaticCredentials("dummydummydummy", "dummydummydummy", "")
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credential,
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("ap-notheast-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   jsonFile,
	})
	if err != nil {
		fmt.Errorf("Upload JSON Error: %s", err)
	}

	log.Printf("Success to upload delay list")
}

/**
 * 鉄道遅延情報提供ページより遅延リストを取得
 */
func getDelayList() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can not get delay list! Error: %v", err)
	}
	defer resp.Body.Close()

	delayList, _ := ioutil.ReadAll(resp.Body)

	return delayList
}

/**
 * 既にS3に上がっている遅延リストを削除
 */
func deleteExistFile() {
	//TODO 既存のJSONを削除する処理の実装
}

/**
 * tempディレクトリにJSONファイルを作成し、遅延リストを書き込む
 */
func createJSON(delayList []byte) error {
	if delayList == nil {
		return fmt.Errorf("create JSON error: %s", "nil bytes was given")
	}

	if err := os.MkdirAll(filepath.Dir(tempDir), 0777); err != nil {
		return err
	}

	file, err := os.Create(tempDir)
	if err != nil {
		return err
	}

	_, err = file.Write(delayList)
	if err != nil {
		return err
	}

	//ファイルの存在確認
	if isExist := isExistTempFile(tempDir); isExist != true {
		return fmt.Errorf("Temp file does not exist")
	}

	return nil
}

func isExistTempFile(tempFile string) bool {
	_, err := os.Stat(tempFile)
	return !os.IsNotExist(err)
}
