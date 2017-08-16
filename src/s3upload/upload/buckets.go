package upload

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

/**
 * Returns a list of buckets
 */
func ListBuckets(svc *s3.S3) []*s3.Bucket {

	result, err := svc.ListBuckets(nil)
	ExitErrorf("Unable to list buckets, %v", err)

	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return result.Buckets
}
