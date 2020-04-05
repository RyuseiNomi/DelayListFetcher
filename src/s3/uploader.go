package s3Uploader

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	bucketName = "delay-list"
	key        = "delay-list.json"
)

/**
 * Yield new session to upload file to S3 bucket
 */
func Upload(jsonFile *os.File) error {

	var sess *session.Session

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

	uploader := s3manager.NewUploader(sess)

	// Upload File With Custom Session
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   jsonFile,
	})
	if err != nil {
		return err
	}

	log.Printf("Succeeded to upload delay list!")
	return nil
}
