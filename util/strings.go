package util

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	RegexValidateStr = "[A-Za-z0-9]+"
)

// Check whether string only contains A-Za-z0-9
func ValidateStr(s string) bool {
	validStr := regexp.MustCompile(RegexValidateStr)
	return validStr.MatchString(s)
}

// TrimSpaceToLower trim spaces and lower string
func TrimSpaceToLower(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

// TrimSpaceToUpper trim spaces and upper string
func TrimSpaceToUpper(str string) string {
	return strings.TrimSpace(strings.ToUpper(str))
}

// append 2 strings
func Append(s1 string, s2 string) string {
	return fmt.Sprintf("%s %s", s1, s2)
}

func NullOrBlankString(s *string) bool {
	if s == nil {
		return true
	}
	return len(strings.TrimSpace(*s)) == 0
}
func EmptyOrBlankString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsEmptyOrBlankString(s string) bool {
	if s == "" || len(strings.TrimSpace(s)) == 0 {
		return true
	}
	return false
}
