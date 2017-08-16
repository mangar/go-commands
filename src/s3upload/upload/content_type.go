package upload

import (
	"regexp"
)

var ContentType = map[string]string{
	"(html|htm)$": "text/html",
	"(js)$":       "application/javascript",
	"(css)$":      "text/css",
	"(gif)$":      "image/gif",
	"(ico)$":      "image/x-icon",
	"(jpg|jpeg)$": "image/jpeg",
	"(json)$":     "application/json",
	"(eot)$":      "application/vnd.ms-fontobject",
	"(otf)$":      "font/otf",
	"(png)$":      "image/png",
	"(pdf)$":      "application/pdf",
	"(zip)$":      "application/zip",
	"(mp4)$":      "video/x-mpeg",
}

/**
 *
 */
func GetContentType(fileName string) string {
	for key, value := range ContentType {
		match, _ := regexp.MatchString(key, fileName)
		if match {
			return value
		}
	}
	return ""
}
