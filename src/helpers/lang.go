package helpers

import (
	"os"
)

func DetectLang() string {
	if golang() {
		return "go"
	}
	return "php"
}

func golang() bool {
	_, err := os.Stat("go.mod")
	if err != nil {
		return false
	}
	return true
}
