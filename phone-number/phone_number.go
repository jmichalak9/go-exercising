// Package phonenumber implements functions to deal with phone numbers
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number returns phone number in standardised format, e.g. "1234567890"
func Number(s string) (string, error) {
	reg, err := regexp.Compile("[^0-9]+")
	cleaned := reg.ReplaceAllString(s, "")

	if len(cleaned) < 10 {
		return "", errors.New("number too short")
	}
	if len(cleaned) > 11 {
		return "", errors.New("number too long")
	}
	if len(cleaned) == 11 {
		if cleaned[0] != '1' {
			return "", errors.New("11 digit number does not start with a 1")
		}
		cleaned = cleaned[1:]
	}
	if int(cleaned[0]-'0') < 2 {
		return "", errors.New("area code starts with number less than 2")
	}
	if int(cleaned[3]-'0') < 2 {
		return "", errors.New("exchange code starts with number less than 2")
	}
	return cleaned, err
}

// AreaCode returns area code from phone number
func AreaCode(s string) (string, error) {
	code, err := Number(s)
	if err != nil {
		return "", err
	}
	return code[0:3], err
}

// Format returns phone number in standardised format, e.g. "(613) 995-0253"
func Format(s string) (string, error) {
	s, err := Number(s)
	if err != nil || len(s) != 10 {
		return "", err
	}
	s = fmt.Sprintf("(%s) %s-%s", s[0:3], s[3:6], s[6:10])
	return s, err
}
