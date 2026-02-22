package checkers

import (
	"fmt"
	"strings"
)

var sensitiveKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"secret",
	"token",
	"api_key",
	"apikey",
	"api-key",
	"auth",
	"credential",
	"private_key",
	"privatekey",
	"private-key",
	"access_token",
	"accesstoken",
	"access-token",
	"refresh_token",
	"refreshtoken",
	"refresh-token",
	"bearer",
	"authorization",
	"session_id",
	"sessionid",
	"session-id",
	"cookie",
	"jwt",
	"ssn",
	"credit_card",
	"creditcard",
	"credit-card",
	"cvv",
	"pin",
}

// CheckNoSensitiveData проверяет сообщение на наличие чувствительных данных,
// используя дефолтный список ключевых слов.
func CheckNoSensitiveData(msg string) error {
	return CheckNoSensitiveDataWithKeywords(msg, nil)
}

// CheckNoSensitiveDataWithKeywords делает то же самое, но позволяет передать
// кастомный список ключевых слов. Если keywords == nil или пустой, берётся
// дефолтный список sensitiveKeywords.
func CheckNoSensitiveDataWithKeywords(msg string, keywords []string) error {
	msgLower := strings.ToLower(msg)

	kw := keywords
	if len(kw) == 0 {
		kw = sensitiveKeywords
	}

	for _, keyword := range kw {
		if strings.Contains(msgLower, keyword) {
			// проверяем, следует ли за ключевым словом ':', '=' или пробел
			idx := strings.Index(msgLower, keyword)
			remaining := msgLower[idx+len(keyword):]

			// паттерны вида "password:", "password=", "password "
			if len(remaining) > 0 {
				nextChar := remaining[0]
				if nextChar == ':' || nextChar == '=' || nextChar == ' ' {
					return fmt.Errorf("contains sensitive keyword '%s' with potential value", keyword)
				}
			}
		}
	}

	return nil
}
