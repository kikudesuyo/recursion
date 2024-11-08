package list

import "strings"

func isValidEmail(str string) bool {
	if strings.Contains(str, " ") {
		return false
	}
	if strings.Count(str, "@") != 1 {
		return false
	}
	atIndex := strings.Index(str, "@")
	if atIndex == 0 {
		return false
	}
	return strings.Index(str[atIndex:], ".") > 0
}

func ValidEmailList(emailList []string) []string {
	validEmails := []string{}
	for _, email := range emailList {
		if isValidEmail(email) {
			validEmails = append(validEmails, email)
		}
	}
	return validEmails
}
