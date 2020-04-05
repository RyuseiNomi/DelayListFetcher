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
	accessKeyID     = os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint        = aws.String("https://s3.ap-northeast-1.amazonaws.com")
)

/**
 * Yield new session to upload file to S3 bucket
 */
func GetUploader() *s3manager.Uploader {
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		endpoint = aws.String("http://172.18.0.2:9000")
	}

	credential := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credential,
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         endpoint,
		S3ForcePathStyle: aws.Bool(true),
	})

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("Load Credential File Error:  %+v\n", err)
	}

	Uploader := s3manager.NewUploader(sess)

	return Uploader
}
