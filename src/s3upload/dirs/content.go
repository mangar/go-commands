/**
 *
 * Get files to be sent to S3 bucket
 *
 */
package dirs

import (
	"os"
	"path/filepath"
)

func GetDirContent(searchDir string) []string {

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	// for _, file := range fileList {
	// 	fmt.Println(file)
	// }

	return fileList
}
