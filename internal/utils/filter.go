package utils

import "strings"

func FilterAllSpace(data string) string {
	return strings.Replace(data, " ", "", -1)
}
func FilterHtml(data string) string {
	data = strings.TrimSpace(data)
	return data
}
