package example

import "log"

func Ignored() {
	msg := "Hello world" // переменная: не строковый литерал в вызове
	log.Println(msg)     // линтер это пропустит
}
