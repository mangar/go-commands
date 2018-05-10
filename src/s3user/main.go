package main

import (
	"fmt"
	"os"
	"s3user/cmd/create"
	"s3user/cmd/defaults"
	"s3user/cmd/delete"
	"s3user/cmd/list"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

// var bucketName = ""

func main() {

	defaults.Key = os.Getenv("AWS_KEY")
	defaults.Secret = os.Getenv("AWS_SECRET")

	if os.Args[1] == "version" {
		printVersion()

	} else if os.Args[1] == "help" {
		fmt.Println("usage: go run s3user create|delete|list USERNAME BUCKET1,BUCKET2,BUCKET3")
		fmt.Println("Ex.:")
		fmt.Println("go run s3user create USERNAME BUCKET1,BUCKET2,BUCKET3")
		fmt.Println("or")
		fmt.Println("go run s3user deletecreate USERNAME")
		fmt.Println("or")
		fmt.Println("go run s3user list")
		fmt.Println("")

	} else if !hasAwsKeySecretSet() {
		fmt.Println("You must specify the AWS_KEY and AWS_SECRET environments variable")

	} else if os.Args[1] == "create" {
		defaults.Username = os.Args[2]
		defaults.BucketName = os.Args[3]
		create.CreateUser()

	} else if os.Args[1] == "delete" {
		defaults.Username = os.Args[2]
		delete.DeleteUser()

	} else if os.Args[1] == "list" {
		list.ListUsers()

	} else {
		fmt.Println("Buy me a beer")

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
