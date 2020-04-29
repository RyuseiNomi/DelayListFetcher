package s3Worker

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	uploadBucket   = "delay-list"
	uploadFileName = "delay-list.json"
)

/**
 * Yield new session to upload file to S3 bucket
 */
func Upload(jsonFile *os.File, sess *session.Session) error {

	_, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("Load Credential File Error:  %+v\n", err)
	}

	uploader := s3manager.NewUploader(sess)

	// Upload File With Custom Session
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(uploadBucket),
		Key:    aws.String(uploadFileName),
		Body:   jsonFile,
	})
	if err != nil {
		return err
	}

	log.Printf("Succeeded to upload delay list!")
	return nil
}
