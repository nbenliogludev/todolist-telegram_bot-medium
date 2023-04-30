package utils

import (
	"strings"
)

func GetKeyValue(str string) (string, string) {
	keyValue := strings.Split(str, "=")
	return keyValue[0], keyValue[1]
}
