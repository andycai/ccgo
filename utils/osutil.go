package ccgo

import (
	"os"
)

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir()
	}

	return false
}
