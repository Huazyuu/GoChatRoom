package server

import (
	"log"
	"testing"
)

func TestLogToDb(t *testing.T) {
	aff := LogToDb("hello", "123")
	if aff != 1 {
		log.Fatal("err")
	}
}
