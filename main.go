package main

import (
	"fmt"
	"os"
)

func main() {
	var name *string

	if len(os.Args) > 1 {
		name = &os.Args[1]
	}

	fmt.Println(greeting(name))
}

func greeting(name *string) string {
	if name != nil && *name != "" {
		return fmt.Sprintf("Hello, %s!", *name)
	}

	return "Hello!"
}
