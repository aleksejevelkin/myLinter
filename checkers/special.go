package checkers

import (
	"fmt"
	"strings"
)

var forbiddenChars = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '~',
}

func CheckSpecialChars(msg string) error {
	// Check for repeated punctuation
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

	// Check for emojis (Unicode ranges for common emojis)
	for _, r := range msg {
		if isEmoji(r) {
			return fmt.Errorf("contains emoji '%c' (U+%04X)", r, r)
		}
	}

	// Check for forbidden special characters (except when used appropriately)
	for _, r := range msg {
		for _, forbidden := range forbiddenChars {
			if r == forbidden {
				// Allow single ! or ? at the end as it might be intentional
				// But multiple !, @, #, etc. are not allowed
				if r != '!' && r != '?' {
					return fmt.Errorf("contains special character '%c'", r)
				}
			}
		}
	}

	// Check for ellipsis patterns
	if strings.Contains(msg, "...") {
		return fmt.Errorf("contains ellipsis '...'")
	}

	return nil
}

// isEmoji checks if a rune is an emoji.
func isEmoji(r rune) bool {
	// Common emoji ranges
	return (r >= 0x1F300 && r <= 0x1F9FF) || // Miscellaneous Symbols and Pictographs, Emoticons, etc.
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
		(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
		(r >= 0x1F1E0 && r <= 0x1F1FF) // Flags
}
