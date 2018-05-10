package main

import (
	"fmt"
	"gmail/cmd/defaults"
	"gmail/cmd/send"
	"os"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

var (
	Version string
	Build   string
	Hash    string
)

var bucketName = ""

func main() {

	defaults.Email = os.Getenv("GMAIL_EMAIL")
	defaults.Pwd = os.Getenv("GMAIL_PWD")

	if !hasGmailSet() {
		fmt.Println("You must specify the GMAIL_EMAIL and GMAIL_PWD")
		os.Exit(3)
	}

	if len(os.Args) == 1 {
		fmt.Println("Buy me a beer")
		os.Exit(3)
	}

	if os.Args[1] == "version" {
		printVersion()

	} else if os.Args[1] == "help" {
		fmt.Println("usage: gmail send TO_EMAIL, SUBJECT, CONTENT")
		fmt.Println("")

	} else if os.Args[1] == "send" {
		defaults.To = os.Args[2]
		defaults.Subject = os.Args[3]
		defaults.Content = os.Args[4]
		send.Send()

	} else {
		fmt.Println("Buy me a beer")

	}

}

func hasGmailSet() bool {
	if defaults.Email == "" || defaults.Pwd == "" {
		return false
	} else {
		return true
	}
}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
