package security

import "golang.org/x/crypto/bcrypt"

type PasswordManager struct{}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{}
}

func (p *PasswordManager) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}

func (p *PasswordManager) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
