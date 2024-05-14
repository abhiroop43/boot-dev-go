package main

import "unicode"

func isValidPassword(password string) bool {
	passwordLen := len(password)
	isUpperCasePresent, isNumberPresent := false, false
	if passwordLen < 5 || passwordLen > 12 {
		return false
	}

	for _, character := range password {
		if unicode.IsUpper(character) {
			isUpperCasePresent = true
		}

		if unicode.IsDigit(character) {
			isNumberPresent = true
		}
	}

	if isNumberPresent && isUpperCasePresent {
		return true
	}

	return false
}
