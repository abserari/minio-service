package minio

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v6"
)

// Upload file
func UploadFile(bucketName string, objectName string, filePath string, contentType string) error {
	endpoint := "127.0.0.1:9001"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	opts := minio.PutObjectOptions{ContentType: contentType}
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, opts)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	return err
}

// dowload file
func DowloadFile(bucketName string, objectName string, filePath string) error {
	endpoint := "127.0.0.1:9001"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"discover.mp3\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully generated presigned URL\n", presignedURL)
	return err
}
