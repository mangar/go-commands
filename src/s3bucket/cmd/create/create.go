package create

import (
	"fmt"
	"s3bucket/cmd/defaults"
	"s3bucket/myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket() {

	if hasBucketName() {
		fmt.Println("Creating Bucket:", defaults.BucketName, " Region:", defaults.Region)
		create()

	} else {
		fmt.Println("You must specify a bucket name.")

	}
}

//
//
func create() {
	if bucketAlreadyCreated() {
		fmt.Println("Bucket already exits")
	} else {
		fmt.Println(" - creating bucket")
		var sess = myaws.GetAwsSession()
		var svc = s3.New(sess)

		_, err := svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(defaults.BucketName),
		})
		if err != nil {
			fmt.Println("Unable to create bucket", defaults.BucketName, err)
		} else {
			setWebsite()
		}
	}

}

//
//
func setWebsite() {
	fmt.Println(" - seting as static website")
	var sess = myaws.GetAwsSession()
	var svc = s3.New(sess)

	params := s3.PutBucketWebsiteInput{
		Bucket: aws.String(defaults.BucketName),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String("index.html"),
			},
			ErrorDocument: &s3.ErrorDocument{
				Key: aws.String("error.html"),
			},
		},
	}

	_, err := svc.PutBucketWebsite(&params)
	if err != nil {
		fmt.Println("Unable to set bucket %q website configuration:", defaults.BucketName, err)
	} else {
		setBucketPolicy()
	}

}

//
//
func setBucketPolicy() {
	fmt.Println(" - applying bucket policies")

	policy := "{\"Version\": \"2012-10-17\", \"Statement\": [{\"Sid\": \"AddPerm\",\"Effect\": \"Allow\",\"Principal\": {\"AWS\": \"*\"},\"Action\": \"s3:GetObject\",\"Resource\": \"arn:aws:s3:::" + defaults.BucketName + "/*\"}]}"

	var sess = myaws.GetAwsSession()
	var svc = s3.New(sess)

	_, err := svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(defaults.BucketName),
		Policy: aws.String(string(policy)),
	})
	if err != nil {
		fmt.Println("Unable to set bucket policy:", defaults.BucketName, err)
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
