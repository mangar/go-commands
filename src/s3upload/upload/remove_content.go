package upload

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var _svc *s3.S3
var _bucketName string

/**
 * Remove all content from the Bucket
 */
func RemoveContent(svc *s3.S3, bucketName string) {

	_svc = svc
	_bucketName = bucketName

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName)})
	ExitErrorf("Unable to list items in bucket %q, %v", err, bucketName)

	num_objs := len(resp.Contents)

	// Create Delete object with slots for the objects to delete
	var objs = make([]*s3.ObjectIdentifier, num_objs)

	for i, o := range resp.Contents {
		objs[i] = &s3.ObjectIdentifier{Key: aws.String(*o.Key)}
		removeObject(objs[i])
	}
}

//
//
func removeObject(o *s3.ObjectIdentifier) {
	_, err := _svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(_bucketName), Key: o.Key})
	ExitErrorf("%v - Unable to delete object %q from bucket %q", err, o, _bucketName)

	fmt.Println("Item removed: ", *o.Key)
}

/**
 *
 */
func ClearBucket(svc *s3.S3, bucketName string) {

	_svc = svc
	_bucketName = bucketName

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName)})
	ExitErrorf("Unable to list items in bucket:", err, bucketName)

	num_objs := len(resp.Contents)

	// Create Delete object with slots for the objects to delete
	var items s3.Delete
	var objs = make([]*s3.ObjectIdentifier, num_objs)

	for i, o := range resp.Contents {
		objs[i] = &s3.ObjectIdentifier{Key: aws.String(*o.Key)}
	}
	// Add list of objects to delete to Delete object
	items.SetObjects(objs)

	svc.DeleteObjects(&s3.DeleteObjectsInput{Bucket: aws.String(_bucketName), Delete: &items})
	// ExitErrorf("Unable to delete objects from bucket %q, %v", err, bucketName)

	fmt.Println("Deleted", num_objs, "object(s) from bucket", bucketName)
}
