package delete

import (
	"fmt"
	"s3bucket/cmd/defaults"
	"s3bucket/myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteBucket() {

	var sess = myaws.GetAwsSession()
	var svc = s3.New(sess)

	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(defaults.BucketName),
	})
	if err != nil {
		fmt.Println("Unable to delete bucket ", defaults.BucketName, err)
	} else {
		fmt.Println("Bucket", defaults.BucketName, "deleted")
	}

}
