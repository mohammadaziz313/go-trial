package greetings

import "fmt"

// Hello is a function
// variable is followed by its type
// It returns a string
func Hello(name string) string{
  // formatting the string
  message:= fmt.Sprintf("Hi, %v. Welcome!",name)
  return message
}
