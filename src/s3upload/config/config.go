/**
 *
 * Read the content of config file in JSON format
 *
 */
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Bucket         string `json:"bucket"`
	AccessKey      string `json:"access_key"`
	Secret         string `json:"secret"`
	Region         string `json:"region"`
	SourceDir      string `json:"source_dir"`
	ReplaceContent bool   `json:"replace_content"`
	DebugMode      bool   `json:"debug"`
}

func ReadConfig(fileLocation string) Config {

	file, e := ioutil.ReadFile(fileLocation)
	if e != nil {
		fmt.Printf("Config file not found: %v\n", e)
		os.Exit(1)
	}

	s := string(file)

	var config Config
	err := json.NewDecoder(strings.NewReader(s)).Decode(&config)

	if err != nil {
		fmt.Printf("Problem decoding config file: %v\n", err)
		os.Exit(1)
	}

	return config
}
