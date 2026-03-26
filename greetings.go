package greetings

import (
	"rsc.io/quote"
)

func Greet() string {
	return "Here, from afiguera greetings repo: " + quote.Hello()
}