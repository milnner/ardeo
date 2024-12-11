package utils

import (
	"regexp"
)

const (
	MAX_PASSWORD_LEN = 32
	MIN_PASSWORD_LEN = 12
)

func ValidateUserPswd(pswd *string) bool {
	// pelo menos um ( @ ou # ou $ ou % ou ! ou & ) e pelo menos um de a..z e de A..Z e de 0..9
	hasSpecialChar := regexp.MustCompile(`.[!@#$%&].`)
	hasUppercase := regexp.MustCompile(`[A-Z]`)
	hasLowercase := regexp.MustCompile(`[a-z]`)
	hasDigit := regexp.MustCompile(`[0-9]`)

	return (hasSpecialChar.MatchString(*pswd) && hasUppercase.MatchString(*pswd) &&
		hasLowercase.MatchString(*pswd) &&
		hasDigit.MatchString(*pswd) &&
		len(*pswd) >= MIN_PASSWORD_LEN &&
		len(*pswd) <= MAX_PASSWORD_LEN)

}

func ValidateUserEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
