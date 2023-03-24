package main

import (
	"fmt"

	"github.com/sergiofgonzalez/quotes"
	"github.com/sergiofgonzalez/quotes/computers"
	"github.com/sergiofgonzalez/quotes/poetry"
	"github.com/sergiofgonzalez/quotes/tv"
)

func main() {

	fmt.Println("=== Random quotes")
	for i := 0; i < 5; i++ {
		fmt.Println(quotes.Quote())
	}

	fmt.Println("\n=== Poetry")
	fmt.Println(poetry.Quote())

	fmt.Println("\n=== Tv")
	fmt.Println(tv.Quote())
	fmt.Println(tv.Quote())

	fmt.Println("\n=== Computers")
	for i := 0; i < 10; i++ {
		fmt.Println(computers.Quote())
	}
}