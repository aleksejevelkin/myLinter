package p

import "log"

func testSpecial() {
	log.Println("something went wrong") // ok
	log.Println("user@domain.com")      // want `log message issue: contains special character`
	log.Println("error #42 occurred")   // want `log message issue: contains special character`
	log.Println("loading...")           // want `log message issue: contains repeated punctuation`
	log.Println("what???")              // want `log message issue: contains repeated punctuation`
	log.Println("request failed")       // ok
}
