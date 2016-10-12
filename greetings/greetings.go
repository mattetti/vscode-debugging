package greetings

import "fmt"

var (
	user string
)

func SayHello(name string) string {
	user = name
	return fmt.Sprintf("Hi %s", user)
}

func SayBye() string {
	return fmt.Sprintf("bye %s", user)
}
