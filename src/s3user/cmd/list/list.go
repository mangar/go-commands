package list

import (
	"fmt"
	"s3user/myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func ListUsers() {
	list()
}

//
//
func list() {
	var sess = myaws.GetAwsSession()
	var svc = iam.New(sess)

	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for i, user := range result.Users {
		if user == nil {
			continue
		}
		fmt.Printf("%d [u] %s\n", i, *user.UserName)
	}

}
