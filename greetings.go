package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Greet() {
	fmt.Println("Here, from afiguera greetings repo: ", quote.Hello())
}