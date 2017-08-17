package main

import (
	"fmt"
	"os"
	"s3upload/config"
	"s3upload/dirs"
	"s3upload/upload"
	"time"
)

var (
	Version string
	Build   string
	Hash    string
)

//
//
func main() {

	configFile := "./config.json"
	if len(os.Args) >= 1 {
		if os.Args[1] == "--version" {
			printVersion()

		} else {
			fmt.Println("[", time.Now(), "] Starting")

			configFile = os.Args[1]

			// 1
			_config := config.ReadConfig(configFile)

			// 2
			files2Upload := dirs.GetDirContent(_config.SourceDir)
			fmt.Println(len(files2Upload), "files do be uploaded.")

			// 3 starts to upload...
			upload.Upload(files2Upload, _config)

			fmt.Println("[", time.Now(), "] Finished")

		}
	}

}

//
//
func printVersion() {
	fmt.Println("Version    :", Version)
	fmt.Println("Build Time :", Build)
	fmt.Println("Hash       :", Hash)
}
