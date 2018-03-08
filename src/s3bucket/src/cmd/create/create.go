package create

import (
	"cmd/defaults"
	"fmt"
	"myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket() {

	if hasBucketName() {
		fmt.Println("Creating Bucket: ", defaults.BucketName, "  Region: ", defaults.Region)
		create()

	} else {
		fmt.Println("You must specify a bucket name.")

	}
}

func create() {
	if bucketAlreadyCreated() {
		fmt.Println("Bucket already exits")
	} else {
		fmt.Println("Creating...")
		var sess = myaws.GetAwsSession()
		var svc = s3.New(sess)

		_, err := svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(defaults.BucketName),
		})
		if err != nil {
			fmt.Println("Unable to create bucket %q, %v", defaults.BucketName, err)
		}

		// // Wait until bucket is created before finishing
		// fmt.Printf("Waiting for bucket %q to be created...\n", bucket)

		// err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		// 	Bucket: aws.String(bucket),
		// })

	}

}

func bucketAlreadyCreated() bool {
	var ret = false

	var sess = myaws.GetAwsSession()
	var svc = s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Unable to list buckets, %v", err)
	}

	// fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		if aws.StringValue(b.Name) == defaults.BucketName {
			ret = true
		}
		// fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	return ret

}

func hasBucketName() bool {
	if defaults.BucketName == "" {
		return false
	} else {
		return true
	}
}
