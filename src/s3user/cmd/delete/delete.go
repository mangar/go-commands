package delete

import (
	"fmt"
	"s3user/cmd/defaults"
	"s3user/myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func DeleteUser() {
	delete()
}

//
//
func delete() {
	var sess = myaws.GetAwsSession()
	var svc = iam.New(sess)

	// err := svc.ListUsers(
	// 	&iam.ListUsersInput{},
	// 	func(page *iam.ListAttachedRolePoliciesOutput, lastPage bool) bool {
	// 		if page != nil && len(page.AttachedPolicies) > 0 {
	// 			for _, policy := range page.AttachedPolicies {
	// 				fmt.Println("Policy:", policy)
	// 			}
	// 			return true
	// 		}
	// 		return false
	// 	},
	// )

	// result, err := svc.DeleteUser(input)

	input := &iam.ListPoliciesInput{
		GroupName: aws.String("Admins"),
	}

	result, err := svc.ListGroupPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

	input := &iam.DeleteUserInput{
		UserName: aws.String(defaults.Username),
	}

	_, err := svc.DeleteUser(input)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	// _, err := svc.DeleteUser(&iam.DeleteUserInput{
	// 	UserName: &defaults.Username,
	// })

	// if err != nil {
	// 	fmt.Println("Problems to remove user:", defaults.Username, err)
	// } else {
	// 	fmt.Println("User deleted.")
	// }

}

// func (c *IAM) ListPolicies(input *ListPoliciesInput) (*ListPoliciesOutput, error)

// 1
// list policies

// 2
// remove policies
