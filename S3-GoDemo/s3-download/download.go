package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// Create an S3 client
	svc := s3.NewFromConfig(cfg)

	// Define the bucket name
	bucket := "s3-godemo" // Replace with your bucket name

	// List and download all objects in the specified S3 bucket
	if err := listAndDownloadObjects(svc, bucket); err != nil {
		log.Fatalf("Failed to list and download objects: %v", err)
	}

	fmt.Println("Successfully downloaded all objects from the bucket.")
}

// listAndDownloadObjects lists and downloads all objects from the specified S3 bucket
func listAndDownloadObjects(svc *s3.Client, bucket string) error {
	output, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return fmt.Errorf("unable to list items in bucket %q, %v", bucket, err)
	}

	for _, object := range output.Contents {
		objectKey := aws.ToString(object.Key)
		log.Printf("Downloading %s", objectKey)

		if err := downloadObject(svc, bucket, objectKey); err != nil {
			return fmt.Errorf("failed to download %q, %v", objectKey, err)
		}
	}

	return nil
}

// downloadObject downloads an object from the specified S3 bucket
func downloadObject(svc *s3.Client, bucket, key string) error {
	output, err := svc.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to download item %q, %v", key, err)
	}
	defer output.Body.Close()

	// Create a local file to save the downloaded object
	filePath := filepath.Join("./downloads", key)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("unable to create directory %q, %v", filepath.Dir(filePath), err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file %q, %v", filePath, err)
	}
	defer file.Close()

	// Write the content of the S3 object to the local file
	_, err = io.Copy(file, output.Body)
	if err != nil {
		return fmt.Errorf("unable to write file %q, %v", filePath, err)
	}

	log.Printf("Successfully downloaded %q to %q", key, filePath)
	return nil
}
