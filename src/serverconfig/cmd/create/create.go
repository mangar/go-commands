package create

import (
	"fmt"
	"net/http"
	"os"
	"serverconfig/cmd/defaults"
)

func CreateBucket() {
	fmt.Println("Creating Config for:", defaults.ProjectName)

	fmt.Println("  - master")
	var serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/master"
	_, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// req, err := http.NewRequest("PUT", serviceUrl, nil)
	// req.Header.Add("version", "30000")
	// _, err := client.Do(req)

	fmt.Println("  - staging")
	serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/staging"
	_, err = http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println("  - develop")
	serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/develop"
	_, err = http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

}
