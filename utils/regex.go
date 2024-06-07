package utils

import (
	"regexp"
	"strings"
)

func StandardizePhone(phone string) string {
	// Regular expression to match various phone number formats
	re := regexp.MustCompile(`^([ \+]*)([0-9\-]{3,})([ \.\-\/]*)([0-9\-]{3,})([ \.\-\/]*)([0-9\-]{4,})$`)

	// Extract matched groups
	match := re.FindStringSubmatch(phone)
	if match == nil {
		return ""
	}

	// Standardize the phone number
	standardized := match[2] + match[4] + match[6]

	return standardized
}

func StandardizeEmail(email string) string {
	// Regular expression to match basic email format
	re := regexp.MustCompile(`^([a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+)@([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*)`)

	// Extract matched groups
	match := re.FindStringSubmatch(email)
	if match == nil {
		return ""
	}

	// Standardize the email
	localPart := strings.ToLower(match[1])  // Convert local part to lowercase
	domainPart := strings.ToLower(match[2]) // Convert domain part to lowercase
	standardized := localPart + "@" + domainPart

	return standardized
}
