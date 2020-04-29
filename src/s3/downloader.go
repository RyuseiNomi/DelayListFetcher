package s3Worker

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	downloadTargetBucket = "delay-list"
	tsvTempDir           = "/tmp/tsv/"
	downloadFileName     = "train.tsv"
)

func DownloadTrainMaster() error {

	var sess *session.Session

	// TODO Credential取得の処理を共通化
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		/* Yield credential for local */
		log.Printf("Start process getting credential as a local")
		credential := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
		sess, _ = session.NewSession(&aws.Config{
			Credentials:      credential,
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String("http://172.18.0.2:9000"),
			S3ForcePathStyle: aws.Bool(true),
		})
	} else {
		/* Yield credential for production */
		log.Printf("Start process getting credential as a production")
		sess, _ = session.NewSession(&aws.Config{
			Region:           aws.String("ap-northeast-1"),
			S3ForcePathStyle: aws.Bool(true),
		})
	}

	_, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("Load Credential File Error:  %+v\n", err)
	}

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
