package util

import (
	"encoding/base64"
	"io/ioutil"

	"strings"
)

const imageFolder string = "../../../static/images/"
const imageURI string = "http://13.238.142.238:5099/images/"

func SaveImage(photoBase64 string, table string, filename string) (string, error) {
	imageContents := strings.Split(photoBase64, ";")
	imageBytes, err := base64.StdEncoding.DecodeString(imageContents[2])
	if err != nil {
		return filename, err
	}

	// Save image to local folder
	filePath := table + "/" + filename + "." + imageContents[1]

	err = ioutil.WriteFile(imageFolder+filePath, imageBytes, 0644)
	if err != nil {
		return filename, err
	}

	return imageURI + filePath, nil
}
