package checkers

import (
	"fmt"
	"strings"
	"unicode"
)

func CheckLowercase(msg string) error {

	if len(msg) == 0 {
		return nil
	}

	// Skip leading whitespace
	trimmed := strings.TrimSpace(msg)
	if len(trimmed) == 0 {
		return nil
	}

	firstRune := []rune(trimmed)[0]

	// If it's a letter, it should be lowercase
	if unicode.IsLetter(firstRune) && unicode.IsUpper(firstRune) {
		return fmt.Errorf("starts with uppercase '%c'", firstRune)
	}

	return nil
}
