package s3utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const (
	S3BucketName = "go-hrms"
)

// UploadToS3 uploads a file to S3 and returns the file URL
func UploadToS3(file *multipart.FileHeader, folder string) (string, error) {
	// Load the AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("your-region"))
	if err != nil {
		return "", fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	// Create an S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file, %v", err)
	}
	defer fileContent.Close()

	// Generate a unique file name
	fileName := fmt.Sprintf("%s-%d-%s", folder, time.Now().UnixNano(), file.Filename)

	// Upload the file to S3
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(fileName),
		Body:        fileContent,
		ContentType: aws.String(file.Header.Get("Content-Type")),
		ACL:         types.ObjectCannedACLPublicRead, // To make the file publicly accessible
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3, %v", err)
	}

	// Return the URL of the uploaded file
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", S3BucketName, cfg.Region, fileName)
	return fileURL, nil
}
