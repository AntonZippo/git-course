package main

import (
	"errors"
	"fmt"
)

func main() {
	var name, language string
	fmt.Println("enter name")
	fmt.Scan(&name)
	fmt.Println("enter language")
	fmt.Scan(&language)
	a, _ := Hello(name, language)
	fmt.Println(a)
}

func Hello(name string, language string) (string, error) {
	if name == "" {
		name = "World"
	}

	prefix := ""

	switch language {
	case "english":
		prefix = "Hello"
	case "spanish":
		prefix = "Hola"
	case "german":
		prefix = "Hallo"
	default:
		return "", errors.New("need to provide a supported language")
	}

	return prefix + " " + name, nil
}
