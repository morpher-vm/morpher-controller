package common

import "strings"

func AwsStr(s string) *string { return &s }
func AwsI32(i int32) *int32   { return &i }
func OrDefault(v, def string) string {
	if strings.TrimSpace(v) == "" {
		return def
	}
	return v
}
