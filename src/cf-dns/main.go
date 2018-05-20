package main

import (
	"cf-dns/cmd/create"
	"cf-dns/cmd/defaults"
	"fmt"
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

	defaults.Email = os.Getenv("CF_EMAIL")
	defaults.Key = os.Getenv("CF_KEY")

	if !hasEnvironmentSet() {
		fmt.Println("You must specify the CF_EMAIL and CF_KEY")
		os.Exit(3)
	}

	if len(os.Args) == 1 {
		fmt.Println("Buy me a beer")
		os.Exit(3)
	}

	if os.Args[1] == "version" {
		printVersion()

	} else if os.Args[1] == "help" {
		fmt.Println("usage: cfdns create DNS_ENTRY, RECORD_TYPE, DNS_VALUE, ZONE")
		fmt.Println("usage: cfdns create www CNAME www-aws-bucket.aws.com zone.com.br")
		fmt.Println("")

	} else if os.Args[1] == "create" {

		if len(os.Args) <= 5 {
			fmt.Println("You have to inform all 3 parameters: DnsName, Record Type, DNS Value and Zone")
			os.Exit(-1)
		}

		defaults.DnsName = os.Args[2]
		defaults.RecordType = os.Args[3]
		defaults.DnsValue = os.Args[4]
		defaults.Zone = os.Args[5]
		create.Create()

	} else {
		fmt.Println("Buy me a beer")

	}

}

func hasEnvironmentSet() bool {
	if defaults.Email == "" || defaults.Key == "" {
		return false
	} else {
		return true
	}
}

func printVersion() {
	out := emoji.Sprint(":checkered_flag: Version: ", Version, "\n:clock2: Build Time: ", Build)
	fmt.Println(out)

}
