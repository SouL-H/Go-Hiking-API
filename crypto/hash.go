package crypto

import (

    "golang.org/x/crypto/bcrypt"
)
//Hashes the incoming password.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
//It checks whether the incoming hash and the password are equal.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
