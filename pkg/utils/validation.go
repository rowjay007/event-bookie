package utils

import (
    "errors"
    "regexp"
)

// ValidateInput validates user input
func ValidateInput(input string) error {
    // Check if input is empty
    if input == "" {
        return errors.New("input is required")
    }

    // Check if input matches a specific pattern, for example, an email address
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(input) {
        return errors.New("invalid email format")
    }

    // Add more validation rules as needed

    return nil
}
