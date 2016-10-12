package main

import (
	"fmt"

	"github.com/mattetti/vscode-debugging/greetings"
)

func main() {
	fmt.Println(greetings.SayHello("Matt"))
	fmt.Println(greetings.SayBye())
}
