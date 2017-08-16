package main

import (
	"fmt"
	"os"
	"s3upload/config"
	"s3upload/dirs"
	"s3upload/upload"
	"time"
)

func main() {
	fmt.Println("[", time.Now(), "] Starting")

	configFile := "./config.json"
	if len(os.Args) >= 1 {
		configFile = os.Args[1]
	}

	// 1
	_config := config.ReadConfig(configFile)

	// 2
	files2Upload := dirs.GetDirContent(_config.SourceDir)
	fmt.Println(len(files2Upload), "files do be uploaded.")

	// 3 starts to upload...
	upload.Upload(files2Upload, _config)

	fmt.Println("[", time.Now(), "] Finished")
}
