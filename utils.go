package main

import (
	"strings"
)

func GetFileName(fileName string) string {
	slicedSlash := strings.Split(fileName, "/")
	slicedBackslash := strings.Split(slicedSlash[len(slicedSlash)-1], "\\")
	return slicedBackslash[len(slicedBackslash)-1]
}
