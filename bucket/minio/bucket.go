package minio

import (
	"log"

	"github.com/minio/minio-go/v6"
)

// // Initialize minio client object
// func InitMinioClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) error {
// 	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
// 	return err
// }

// Make a new bucket
func MakeBucket(bucketName string, location string) error {
	endpoint := "127.0.0.1:9001"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	return err
}
