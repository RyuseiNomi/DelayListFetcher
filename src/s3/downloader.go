package s3Worker

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	downloadTargetBucket = "delay-list"
	tsvTempDir           = "/tmp/tsv/"
	downloadFileName     = "train.tsv"
)

func DownloadTrainMaster(sess *session.Session) error {

	if err := os.MkdirAll(tsvTempDir, 0777); err != nil {
		return err
	}
	file, err := os.Create(tsvTempDir + downloadFileName)
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	tsvFile, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(downloadTargetBucket),
		Key:    aws.String(downloadFileName),
	})
	if err != nil {
		return err
	}

	log.Printf("Succeeded to download train master!: %d byte", tsvFile)
	return nil
}
