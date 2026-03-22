package validator

import (
	"errors"
	"regexp"

	"jatinporiya/models"
)

var ValidTypes = map[models.CtaType]struct{}{
	models.BookDemo:    {},
	models.TalkToSales: {},
	models.Career:      {},
	models.Contact:     {},
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var phoneRegex = regexp.MustCompile(`^[0-9]{10}$`)

func ValidateCta(cta models.Cta) error {

	// Type validation
	for _, t := range cta.Type {
		if _, ok := ValidTypes[t]; !ok {
			return errors.New("Invalid type!")
		}
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
