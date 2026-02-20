package checkers

import (
	"fmt"
	"strings"
)

var forbiddenChars = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '~',
}

func CheckSpecialChars(msg string) error {
	// проверка повторяющейся пунктуации
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

	// проверка эмодзи
	for _, r := range msg {
		if isEmoji(r) {
			return fmt.Errorf("contains emoji '%c' (U+%04X)", r, r)
		}
	}

	// проверка запрещённых спецсимволов
	for _, r := range msg {
		for _, forbidden := range forbiddenChars {
			if r == forbidden {
				if r != '!' && r != '?' {
					return fmt.Errorf("contains special character '%c'", r)
				}
			}
		}
	}

	// проверка многоточия
	if strings.Contains(msg, "...") {
		return fmt.Errorf("contains ellipsis '...'")
	}

	return nil
}

func isEmoji(r rune) bool {
	// Распространённые диапазоны эмодзи
	return (r >= 0x1F300 && r <= 0x1F9FF) || // Разные символы и пиктограммы, эмотиконы и т.д.
		(r >= 0x2600 && r <= 0x26FF) || // Разные символы
		(r >= 0x2700 && r <= 0x27BF) || // Дингбаты
		(r >= 0x1F600 && r <= 0x1F64F) || // Эмотиконы
		(r >= 0x1F680 && r <= 0x1F6FF) || // Символы транспорта и карт
		(r >= 0x1F1E0 && r <= 0x1F1FF) // Флаги
}
