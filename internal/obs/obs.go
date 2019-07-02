package obs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

type OBS struct {
	BucketName string
	S3Uploader *s3manager.Uploader
	S3         *s3.S3
}

func New(c m.ConfigOBS) *OBS {
	obs := OBS{}
	obs.BucketName = c.BucketName
	s3session := createAWSSession(c)
	obs.S3Uploader = createUploader(s3session)
	obs.S3 = createS3(s3session)
	return &obs
}

func createAWSSession(c m.ConfigOBS) *session.Session {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AccessKeyID, c.SecretAccessKey, ""),
		Endpoint:    aws.String(c.Endpoint),
		Region:      aws.String(c.Region),
	}
	newSession := session.New(s3Config)
	return newSession
}

func createUploader(sess *session.Session) *s3manager.Uploader {
	uploader := s3manager.NewUploader(sess)
	return uploader
}

func createS3(sess *session.Session) *s3.S3 {
	svc := s3.New(sess)
	return svc
}

// Upload upload file
func (obs *OBS) Upload(f *os.File) error {
	_, err := obs.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(obs.BucketName),
		Key:    aws.String("temp"),
		Body:   f,
	})
	return err
}

func (obs *OBS) GenPresign() int {
	return 1
}
