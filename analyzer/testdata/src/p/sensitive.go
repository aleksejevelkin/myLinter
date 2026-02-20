package p

import "log"

func testSensitive() {
	log.Println("user logged in")          // ок: нет чувствительных данных
	log.Println("password: 12345")         // want `log message issue: contains sensitive keyword`
	log.Println("token: abc123")           // want `log message issue: contains sensitive keyword`
	log.Println("api_key=secret_value")    // want `log message issue: contains sensitive keyword`
	log.Println("connection established")  // ок: нет чувствительных данных
	log.Println("authorization: Bearer x") // want `log message issue: contains sensitive keyword`
}
