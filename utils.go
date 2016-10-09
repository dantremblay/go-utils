package utils

import (
	"log"
	"strings"
)

func RecoverFunc() {
	if r := recover(); r != nil {
		log.Println("Recovered:", r)
	}
}

func ConvertSliceToMap(sep string, slice []string) map[string]string {
	result := make(map[string]string)

	if len(slice) == 0 {
		return nil
	}

	for _, s := range slice {
		split := strings.Split(s, sep)

		result[split[0]] = split[1]
	}

	return result
}

func StringInSlice(a string, list []string) bool {
	for _, v := range list {
		if a == v {
			return true
		}
	}

	return false
}
