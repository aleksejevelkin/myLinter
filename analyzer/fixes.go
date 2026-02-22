package analyzer

import (
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

func canFixStringLiteral(litValue string) bool {
	// чиним только "..."
	return len(litValue) >= 2 && litValue[0] == '"' && litValue[len(litValue)-1] == '"'
}

func buildReplaceWholeLiteralFix(litPos token.Pos, litValue string, newMsg string) (*analysis.SuggestedFix, bool) {
	if !canFixStringLiteral(litValue) {
		return nil, false
	}
	quoted := strconv.Quote(newMsg)
	return &analysis.SuggestedFix{
		Message: "исправить сообщение",
		TextEdits: []analysis.TextEdit{{
			Pos:     litPos,
			End:     litPos + token.Pos(len(litValue)),
			NewText: []byte(quoted),
		}},
	}, true
}

func fixLowercaseStart(msg string) (string, bool) {
	if len(msg) == 0 {
		return msg, false
	}

	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsUpper(r) {
		lower := unicode.ToLower(r)
		_, size := utf8.DecodeRuneInString(msg)
		return string(lower) + msg[size:], true
	}
	return msg, false
}

func fixSpecial(msg string) (string, bool) {
	changed := false
	out := make([]rune, 0, len([]rune(msg)))

	// убираем часть символов
	for _, r := range msg {
		switch r {
		case '@', '#', '$', '%', '^', '&', '*', '~':
			changed = true
			continue
		}
		out = append(out, r)
	}

	msg2 := string(out)

	// схлопываем повторы ! ? .
	var out2 []rune
	prev := rune(0)
	for _, r := range msg2 {
		if (r == '!' || r == '?' || r == '.') && r == prev {
			changed = true
			continue
		}
		out2 = append(out2, r)
		prev = r
	}
	msg3 := string(out2)

	if msg3 != msg {
		changed = true
	}
	return msg3, changed
}
