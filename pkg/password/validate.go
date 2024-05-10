package password

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(plain string, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
		return false
	}

	return true
}
