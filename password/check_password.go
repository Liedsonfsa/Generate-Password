package password

import "strings"

// CheckPassword check if the password is secure
func CheckPassword(password string) bool {
	if hasUpperCase(password) && hasLowerCase(password) && hasNumber(password) && hasEspecials(password) {
		return true
	}

	return false
}

func hasUpperCase(password string) bool {
	var upper string = "ZXCVBNMLSKDJFHGPOQWIEURYT"
	return strings.ContainsAny(password, upper)
}

func hasLowerCase(password string) bool {
	var lower string = "aqzwsxcderfvbgtyhnmjuiklop"
	return strings.ContainsAny(password, lower)
}

func hasNumber(password string) bool {
	var numbers string = "0123456789"
	return strings.ContainsAny(password, numbers)
}

func hasEspecials(password string) bool {
	var especials string = "!@#$%^&*()_+{}[]:;<>,.?/~`"
	return strings.ContainsAny(password, especials)
}