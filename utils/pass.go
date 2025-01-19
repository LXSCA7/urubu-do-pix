package utils

import (
	// "errors"
	"strings"
	"unicode"
)

func IsPasswordStrong(password string) []string {
	var passErrors []string

	if strings.TrimSpace(password) == "" {
		passErrors = append(passErrors, "a senha nao pode ser vazia")
		return passErrors
	}

	if len(password) < 8 {
		passErrors = append(passErrors, "a senha deve possuir no minimo 8 caracteres")
	}

	if password == "12345678" {
		passErrors = append(passErrors, "12345678 ta de sacanagem ne paizao kkkkkkkkkkkkkkk urubu da seguranca")
	}

	hasSpecial := false
	hasNumber := false
	// hasLetter := false
	hasUpper := false
	hasLower := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			// hasLetter = true
			if unicode.IsUpper(char) {
				hasUpper = true
			} else {
				hasLower = true
			}
		} else if unicode.IsNumber(char) {
			hasNumber = true
		} else {
			hasSpecial = true
		}
	}

	if !hasSpecial {
		passErrors = append(passErrors, "a senha deve possuir um caractere especial")
	}

	if !hasNumber {
		passErrors = append(passErrors, "a senha deve possuir um numero")
	}

	if !hasUpper {
		passErrors = append(passErrors, "a senha deve possuir uma letra maiuscula")
	}

	if !hasLower {
		passErrors = append(passErrors, "a senha deve possuir uma letra minuscula")
	}

	if len(passErrors) == 0 {
		return nil
	}

	return passErrors
}
