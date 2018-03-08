package myaws

import (
	"cmd/defaults"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAwsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(defaults.Region),
		Credentials: credentials.NewStaticCredentials(defaults.Key, defaults.Secret, ""),
	})

	if err != nil {
		fmt.Println("Problems to connect to S3", err)
	}

	return sess
}
