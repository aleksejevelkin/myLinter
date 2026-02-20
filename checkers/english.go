package checkers

import (
	"fmt"
	"unicode"
)

func CheckEnglishOnly(msg string) error {
	for _, r := range msg {
		// разрешаем базовые ASCII-символы (от пробела до тильды)
		if r >= 0x20 && r <= 0x7E {
			continue
		}
		// разрешаем перевод строки и табуляцию
		if r == '\n' || r == '\t' || r == '\r' {
			continue
		}
		// любой другой символ — не английский
		if unicode.IsLetter(r) && r > 127 {
			return fmt.Errorf("contains non-English character '%c' (U+%04X)", r, r)
		}
		// эмодзи или другие юникод-символы
		if r > 127 {
			return fmt.Errorf("contains non-ASCII character '%c' (U+%04X)", r, r)
		}
	}
	return nil
}
