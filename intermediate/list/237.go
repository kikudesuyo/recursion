package list

func isValidEmail(str string) bool {
	runes := []rune(str)
	if string(runes[0]) == "@" {
		return false
	}
	atMarKCnt := 0
	isDot := false
	for _, r := range runes {
		if string(r) == " " {
			return false
		}
		if string(r) == "@" {
			atMarKCnt++
		}
		if atMarKCnt > 1 {
			return false
		}
		if atMarKCnt == 1 && string(r) == "." {
			isDot = true
		}
	}
	return isDot
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
