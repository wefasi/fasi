package storage

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	bucket string
	client *s3.Client
}

func NewS3Storage() S3Storage {
	profile := "fasi"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	return S3Storage{
		bucket: "fasi-sites",
		client: client,
	}
}

func (s *S3Storage) Get(file string) (string, error) {
	output, err := s.client.GetObject(context.TODO(),
		&s3.GetObjectInput{
			Bucket: &s.bucket,
			Key:    &file,
		},
	)
	if err != nil {
		var he *awshttp.ResponseError
		if errors.As(err, &he) {
			if he.HTTPStatusCode() == 404 {
				return "", errors.New("not found")
			}
		}
		return "", err
	}
	defer output.Body.Close()
	body, err := io.ReadAll(output.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil
}

func (s *S3Storage) Put(file string, content string) error {
	buff := strings.NewReader(content)
	_, err := s.client.PutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: &s.bucket,
			Key:    &file,
			Body:   buff,
		},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
