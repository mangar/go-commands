package delete

import (
	"fmt"
	"net/http"
	"os"
	"serverconfig/cmd/defaults"
)

func DeleteBucket() {
	fmt.Println("Deleting Config for:", defaults.ProjectName)
	delete("master")
	delete("staging")
	delete("develop")
}

func delete(environment string) {
	fmt.Println("  -", environment)
	var serviceUrl = defaults.ServerConfigUrl + "/api/v1/version/" + defaults.ProjectName + "/" + environment

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", serviceUrl, nil)
	_, err = client.Do(req)

	// _, err := http.Get(serviceUrl)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}
