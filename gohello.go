package gohello

import "fmt"

// Message return the string "Hello!"
func Message() string {
	return "Hello!"
}

// Say println the Message()
func Say() {
	fmt.Println(Message())
}
