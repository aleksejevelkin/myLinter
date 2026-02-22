package checkers

import (
	"fmt"
	"strings"
)

var forbiddenChars = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '~',
}

func CheckSpecialChars(msg string) error {
	// повторы !!! ??? ...
	prevChar := rune(0)
	repeatCount := 0
	for _, r := range msg {
		if r == prevChar && (r == '!' || r == '?' || r == '.') {
			repeatCount++
			if repeatCount >= 2 {
				return fmt.Errorf("contains repeated punctuation '%c%c%c'", r, r, r)
			}
		} else {
			repeatCount = 1
		}
		prevChar = r
	}

	// эмодзи
	for _, r := range msg {
		if isEmoji(r) {
			return fmt.Errorf("contains emoji '%c' (U+%04X)", r, r)
		}
	}

	// запрещённые символы
	for _, r := range msg {
		for _, forbidden := range forbiddenChars {
			if r == forbidden {
				if r != '!' && r != '?' {
					return fmt.Errorf("contains special character '%c'", r)
				}
			}
		}
	}

	if strings.Contains(msg, "...") {
		return fmt.Errorf("contains ellipsis '...'")
	}

	return nil
}

func isEmoji(r rune) bool {
	return (r >= 0x1F300 && r <= 0x1F9FF) ||
		(r >= 0x2600 && r <= 0x26FF) ||
		(r >= 0x2700 && r <= 0x27BF) ||
		(r >= 0x1F600 && r <= 0x1F64F) ||
		(r >= 0x1F680 && r <= 0x1F6FF) ||
		(r >= 0x1F1E0 && r <= 0x1F1FF)
}
