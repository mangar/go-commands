package upload

import (
	"fmt"
	"regexp"
)

var ContentType = map[string]string{
	"(.html|.htm)$":       "text/html",
	"(.css)$":             "text/css",
	"(.txt)$":             "text/plain",
	"(.gif)$":             "image/gif",
	"(.ico)$":             "image/x-icon",
	"(.jpg|jpeg)$":        "image/jpeg",
	"(.png)$":             "image/png",
	"(.svg)$":             "image/svg+xml",
	"(.js|.map|.js.map)$": "application/javascript",
	"(.eot)$":             "application/vnd.ms-fontobject",
	"(.pdf)$":             "application/pdf",
	"(.zip)$":             "application/zip",
	"(.json)$":            "application/json",
	"(.otf)$":             "font/otf",
	"(.mp4)$":             "video/x-mpeg",
	"(.ogg)$":             "audio/ogg",
	"(.mpeg)$":            "audio/mpeg",
	"(.webm)$":            "audio/webm",
	"(.ttf)$":             "application/x-font-ttf",
}

/**
 *
 */
func GetContentType(fileName string) string {
	fmt.Println("[ContentType] FileName", fileName)
	for key, value := range ContentType {
		match, _ := regexp.MatchString(key, fileName)
		if match {
			fmt.Println("  >>", value)
			return value
		}
	}
	return "not_found"
}
