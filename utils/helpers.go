package utils

import (
	"strings"
	"time"
)

func GetTimestamp() string {
	// get the current date and time
	return time.Now().Format("2006-01-02 15:04:05")
}

func DecodeID(encodedID string) []string {
	return strings.Split(encodedID, "_")

}

func Normalize(input string) string {
	lower := strings.ToLower(input)
	return strings.ReplaceAll(lower, " ", "-")
}
