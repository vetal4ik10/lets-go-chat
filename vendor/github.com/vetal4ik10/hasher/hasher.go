package hasher

import "golang.org/x/crypto/bcrypt"

// HashPassword returns the hash of the password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}