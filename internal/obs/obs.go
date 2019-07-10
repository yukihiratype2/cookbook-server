package obs

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
	"github.com/yukihiratype2/cookbook-server/internal/util"
)

type OBS struct {
	BucketName string
	SESS       *session.Session
	S3         *s3.S3
}

func New() *OBS {
	c := (config.LoadConfig()).OBS
	obs := OBS{}
	obs.BucketName = c.BucketName
	obs.SESS = createAWSSession(c)
	// obs.S3Uploader = createUploader()
	obs.S3 = obs.createS3()
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

func (obs *OBS) createS3() *s3.S3 {
	svc := s3.New(obs.SESS)
	return svc
}

func (obs *OBS) GenPresign(SHA2 string) (urlStr string, err error) {

	resp, _ := obs.S3.PutObjectRequest(&s3.PutObjectInput{
		// Bucket:      aws.String(obs.BucketName),
		Bucket:      aws.String("cookbook"),
		Key:         aws.String(SHA2),
		ContentType: aws.String("image/jpg"),
		// ACL:    aws.String("public-read-write"),
	})
	urlStr, err = resp.Presign(10 * time.Minute)

	return
}
