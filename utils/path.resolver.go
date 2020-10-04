package utils

import (
	"strconv"
	"strings"
)

func ConfirmStringPathVariable(path string, urlpathvar string) string {
	arr := strings.Split(path, "/")

	if 1 < len(arr) {
		if arr[2] == urlpathvar {
			return arr[2]
		}
		return ""
	} else {
		return ""
	}
}

//ConfirmIntPathVariable ..
func ConfirmIntPathVariable(path string) (bool, int) {
	arr := strings.Split(path, "/")

	if 3 < len(arr) {
		if val, err := strconv.Atoi(arr[3]); err != nil {
			return false, -1
		} else {
			return true, val
		}

	} else {
		return false, -1
	}

}
