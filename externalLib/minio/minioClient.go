package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

/*
*
OSS
go client doc: https://min.io/docs/minio/linux/developers/go/minio-go.html#
github：https://github.com/minio/minio
*/

const (
	endpoint        = "127.0.0.1:50021"
	accessKeyID     = "5WXG4qQaYJGDGec3z7hk"
	secretAccessKey = "NtBJ9z2Zb241vl3EjNtJGfsrsiXwgEICeGAb0NDn"
	useSSL          = false
)

var minioClient *minio.Client

func init() {
	_ = os.Chdir("externalLib/minio")

	var err error
	// Initialize minio client object.
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("err ---->  ", err)
	}
	log.Printf("connect sucessfully!\n%#v\n", minioClient) // minioClient is now set up
}

func upload() {
	ctx := context.Background()

	// Make a new bucket called mymusic.
	bucketName := "mydoc"
	location := "us-east-1"

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "img-desktop-1.jpg"
	filePath := "../../assets/img-desktop-1.jpg" // 相对于此路径
	contentType := "image/jpg"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

}

func main() {
	upload()
}
