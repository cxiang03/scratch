package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//nolint:gosec // test
const (
	endpoint  = "proxy.uss.s3.test.shopee.io"
	accessKey = "60086901"
	secretKey = "jTfWrNlmQOJbTDAaqIyKSkAkRGKSvtxH"
)

func main() {
	s3, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	policy, err := s3.GetBucketPolicy(context.Background(), "scratch")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(policy)
}
