package validator

import (
	"errors"
	"regexp"

	"jatinporiya/models"
)

var ValidTypes = map[string]bool{
	"book_demo":     true,
	"talk_to_sales": true,
	"career":        true,
	"contact":       true,
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var phoneRegex = regexp.MustCompile(`^[0-9]{10}$`)

func ValidateCta(cta models.Cta) error {

	// Type validation
	if _, ok := ValidTypes[cta.Type]; !ok {
		return errors.New("Invalid type. Allowed: book_demo, talk_to_sales, career, contact")
	}

	// Basic Validation
	if cta.Name == "" || cta.Email == "" || cta.Phone == "" {
		return errors.New("Input Data is Empty!")
	}

	if !emailRegex.MatchString(cta.Email) {
		return errors.New("Invalid email format")
	}

	if !phoneRegex.MatchString(cta.Phone) {
		return errors.New("Invalid phone number")
	}

	return nil
}
