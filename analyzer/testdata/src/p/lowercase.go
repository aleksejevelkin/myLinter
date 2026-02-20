package p

import "log"

func testLowercase() {
	log.Println("everything is fine")      // ок: строчная буква
	log.Println("Hello world")             // want `log message issue: starts with uppercase 'H'`
	log.Printf("Error occurred", "detail") // want `log message issue: starts with uppercase 'E'`
	log.Println("starting process")        // ок: строчная буква
	log.Println("Started process")         // want `log message issue: starts with uppercase 'S'`
}
