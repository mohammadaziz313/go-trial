package main

import(
  "fmt"
  "test.com/greetings"
)

func main(){
  message := greetings.Hello("Header")
  fmt.Println(message)
}
