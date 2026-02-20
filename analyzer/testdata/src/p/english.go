package p

import "log"

func testEnglish() {
	log.Println("connection timeout")      // ок: только английские символы
	log.Println("привет мир")              // want `log message issue: contains non-English character`
	log.Println("ошибка подключения")      // want `log message issue: contains non-English character`
	log.Println("error: timeout")          // ок: только английские символы
	log.Println("смешанный mixed message") // want `log message issue: contains non-English character`
}
