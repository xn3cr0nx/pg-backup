package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// NewS3Session creates a new S3 session
func NewS3Session(region string) (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// UploadFile uploads a file to s3
func UploadFile(s *session.Session, bucket, filename string, file io.Reader) (err error) {
	uploader := s3manager.NewUploader(s)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	return
}
