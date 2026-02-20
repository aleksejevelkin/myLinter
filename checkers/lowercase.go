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

	// пропускаем начальные пробелы
	trimmed := strings.TrimSpace(msg)
	if len(trimmed) == 0 {
		return nil
	}

	firstRune := []rune(trimmed)[0]

	// если первый символ — буква, она должна быть строчной
	if unicode.IsLetter(firstRune) && unicode.IsUpper(firstRune) {
		return fmt.Errorf("starts with uppercase '%c'", firstRune)
	}

	return nil
}
