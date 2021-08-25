package utils

import "os"

func IsTest() bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	} else {
		return false
	}
}
