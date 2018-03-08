package main

import (
	"fmt"
	"os"
	"serverconfig/cmd/create"
	"serverconfig/cmd/defaults"
	"serverconfig/cmd/delete"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

func main() {
	defaults.ProjectName = os.Args[2]
	defaults.ServerConfigUrl = os.Args[3]

	if os.Args[1] == "version" {
		printVersion()

	} else if os.Args[1] == "help" {
		fmt.Println("usage: serverconfig create|delete PROJECTNAME SERVER_CONFIG_URL")
		fmt.Println("Ex.:")
		fmt.Println("serverconfig create Project1 https://heroku.com/project/server/config")
		fmt.Println("")

	} else if os.Args[1] == "create" {

		create.CreateBucket()

	} else if os.Args[1] == "delete" {
		delete.DeleteBucket()

	} else {
		fmt.Println("Buy")

	}

}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
