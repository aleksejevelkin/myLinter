package checkers

import (
	"fmt"
	"unicode"
)

func CheckEnglishOnly(msg string) error {
	for _, r := range msg {
		// Allow basic ASCII printable characters (space to tilde)
		if r >= 0x20 && r <= 0x7E {
			continue
		}
		// Allow newline and tab
		if r == '\n' || r == '\t' || r == '\r' {
			continue
		}
		// Any other character is non-English
		//if unicode.IsLetter(r) && r > 127 {

		if !unicode.Is(unicode.Latin, r) {
			return fmt.Errorf("contains non-English character '%c' (U+%04X)", r, r)
		}
		// Emoji or other unicode
		if r > 127 {
			return fmt.Errorf("contains non-ASCII character '%c' (U+%04X)", r, r)
		}
	}
	return nil
}
