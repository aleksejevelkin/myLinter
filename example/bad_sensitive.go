package example

import "log"

func BadSensitive() {
	log.Println("password: 12345")
	log.Println("token: abc123")
	log.Println("api_key=secret")
}
