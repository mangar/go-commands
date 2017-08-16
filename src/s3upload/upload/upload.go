package upload

import (
	"fmt"
	"os"
	"s3upload/config"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var _uploader *s3manager.Uploader
var _config *config.Config
var _totalFilesCount int

/*
 * Upload file to S3
 */
func Upload(files []string, config config.Config) {

	_totalFilesCount = len(files)
	_config = &config

	if _config.DebugMode {
		fmt.Println("Bucket:", _config.Bucket)
		fmt.Println("AccessKey:", _config.AccessKey)
		fmt.Println("Secret:", _config.Secret)
		fmt.Println("Region:", _config.Region)
	}

	sess := getAwsSession()
	svc := s3.New(sess)

	if config.ReplaceContent {
		ClearBucket(svc, config.Bucket)
	}

	//
	var control sync.WaitGroup
	control.Add(_totalFilesCount)

	_uploader = s3manager.NewUploader(sess)

	for index, file := range files {
		go uploadFile(index, file, &control)
	}
	control.Wait()
}

//
//
func uploadFile(index int, fileName string, control *sync.WaitGroup) {

	newFileName := strings.Replace(fileName, _config.SourceDir, "", -1)
	fmt.Println("[", index, "/", _totalFilesCount, "] uploading...", newFileName)

	file, err := os.Open(fileName)
	ExitErrorf("Unable to open file %q, %v", err)

	_, err = _uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(_bucketName),
		Key:         aws.String(newFileName),
		Body:        file,
		ContentType: aws.String(GetContentType(newFileName)),
	})

	defer file.Close()

	control.Done()
}

//
//
//
func getAwsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(_config.Region),
		Credentials: credentials.NewStaticCredentials(_config.AccessKey, _config.Secret, ""),
	})
	ExitErrorf("Problems to connect to S3", err)
	return sess
}
