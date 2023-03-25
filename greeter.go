package greeter

import "fmt"

// Hello returns "Hello, World!"
func Hello() string {
	return "Hello, World!"
}


// HelloTo returns a greeting for the given name
func HelloTo(name string) string {
	if len(name) == 0 {
		name = "stranger"
	}
	return fmt.Sprintf("Hello to %s", name)
}