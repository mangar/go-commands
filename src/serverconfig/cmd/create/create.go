package create

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"serverconfig/cmd/defaults"
)

func CreateBucket() {
	fmt.Println("Creating Config for:", defaults.ProjectName)
	create("master")
	create("staging")
	create("develop")
}

func create(environment string) {
	fmt.Println("  -", environment)
	var serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/" + environment
	_, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	updateVersion(environment)
}

func updateVersion(environment string) {
	fmt.Println("    - updatingVersion", environment)
	var serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/" + environment

	client := &http.Client{}
	req, err := http.NewRequest("PUT", serviceUrl, nil)

	form := url.Values{}
	form.Add("version", "10000")
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err = client.Do(req)

	// _, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}
