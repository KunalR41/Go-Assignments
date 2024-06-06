package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	bucketName := "s3-godemo"
	keyName := "example.zip"
	filePath := "C:\\Users\\DELL\\Go Project\\s3-upload\\example.zip"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config,%v,err")
	}
	//create S3 client
	svc := s3.NewFromConfig(cfg)

	//open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to Open file %q,%v", filePath, err)

	}
	defer file.Close()

	//upload file to S3
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Unable to uupload %q to %q,%v", filePath, bucketName, err)
	}
}
