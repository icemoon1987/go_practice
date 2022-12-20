package main

import "fmt"
import "rsc.io/quote"
import "example.com/greetings"

func main() {

	fmt.Println("Hello World!")
	fmt.Println(quote.Go())

	var message string
	message = greetings.Hello("icemoon1987")
	fmt.Println(message)
}
