package storage

import (
	"context"
	"errors"
	"io"
	"log"
	"path"
	"strings"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	bucket string
	folder string
	client *s3.Client
}

func NewS3Storage(folder string) S3Storage {
	profile := "fasi"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	return S3Storage{
		bucket: "fasi-sites",
		folder: folder,
		client: client,
	}
}

func (s *S3Storage) Folder() string {
	return path.Join(s.bucket, s.folder)
}

func (s *S3Storage) SetFolder(folder string) {
	s.folder = folder
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
				// TODO custom errors
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
	key := path.Join(s.folder, file)
	_, err := s.client.PutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: &s.bucket,
			Key:    &key,
			Body:   buff,
		},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
