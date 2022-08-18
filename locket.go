package locketgo

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
)

func getLogs(logDir string) ([]os.FileInfo, error) {
	files, readErr := ioutil.ReadDir(logDir)
	if readErr != nil {
		return nil, readErr
	}

	return files, nil
}

func getLog(logDir string) (*string, error) {
	path := path.Join(logDir, "locket.log")
	fileData, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	fileString := base64.StdEncoding.EncodeToString(fileData[:])

	return &fileString, nil
}
