package dto

import (
	"regexp"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

const (
	passwordRegex = `^[A-Za-z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*[A-Z][A-Za-z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?][A-Za-z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]*$`
	emailRegex    = `^[a-zA-Z0-9._%+-]+@(gmail\.com|astanait\.edu\.kz)$`
)

func validateClientCreateRequest(req ClientCreateRequest) error {
	validations := []func() error{
		func() error { return validateEmail(req.Email) },
		func() error { return validatePassword(req.Password) },
	}

	for _, v := range validations {
		err := v()
		if err != nil {
			return err
		}
	}

	return nil
}

func validateEmail(email string) error {
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return model.ErrInvalidEmail
	}

	return nil
}

func validatePassword(password string) error {
	re := regexp.MustCompile(passwordRegex)
	if !re.MatchString(password) {
		return model.ErrInvalidPassword
	}

	return nil
}
