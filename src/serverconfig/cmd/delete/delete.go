package delete

import (
	"fmt"
	"net/http"
	"os"
	"serverconfig/cmd/defaults"
)

func DeleteBucket() {

	fmt.Println("Deleting Config for:", defaults.ProjectName)

	fmt.Println("  - master")
	var serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/master"

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", serviceUrl, nil)
	_, err = client.Do(req)

	// _, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	//
	//
	fmt.Println("  - staging")
	serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/staging"

	client = &http.Client{}
	req, err = http.NewRequest("DELETE", serviceUrl, nil)
	_, err = client.Do(req)

	// _, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	//
	//
	fmt.Println("  - develop")
	serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/develop"

	client = &http.Client{}
	req, err = http.NewRequest("DELETE", serviceUrl, nil)
	_, err = client.Do(req)

	// _, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

}
