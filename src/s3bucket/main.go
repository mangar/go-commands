package main

import (
	"fmt"
	"os"
	"s3bucket/cmd/create"
	"s3bucket/cmd/defaults"
	"s3bucket/cmd/delete"
	"s3bucket/cmd/list"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

var bucketName = ""

func main() {

	defaults.Key = os.Getenv("AWS_KEY")
	defaults.Secret = os.Getenv("AWS_SECRET")

	if os.Args[1] == "version" {
		printVersion()

	} else if os.Args[1] == "help" {
		fmt.Println("usage: go run s3bucket create|delete BUCKET_NAME")
		fmt.Println("Ex.:")
		fmt.Println("go run s3bucket create YOUR_PUBLIC_BUCKET_NAME AWS_KEY AWS_SECRET")
		fmt.Println("")

	} else if !hasAwsKeySecretSet() {
		fmt.Println("You must specify the AWS_KEY and AWS_SECRET environments variable")

	} else if os.Args[1] == "create" {
		defaults.BucketName = os.Args[2]
		create.CreateBucket()

	} else if os.Args[1] == "delete" {
		defaults.BucketName = os.Args[2]
		delete.DeleteBucket()

	} else if os.Args[1] == "list" {
		list.ListBuckets()

	} else {
		fmt.Println("Buy")

	}

}

func hasAwsKeySecretSet() bool {
	if defaults.Key == "" || defaults.Secret == "" {
		return false
	} else {
		return true
	}
}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
