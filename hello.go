package main

import "fmt"

uageconst espanhol = "espanol"
const french = "french"
const eng = "eng"
const prefixPortOi = "Oi "
const prefixEspanhol = "Hola, "
const prefixFrench = "Bonjour, "
const prefixEng = "Hi, "

func Hello(name, language string) string {
	if name == "" {
		name = "Any"
	}

	prefix := prefixPortOi

	switch language {
	case french:
		prefix = prefixFrench
	case espanhol:
		prefix = prefixEspanhol
	case eng:
		prefix = prefixEng
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("anyone", "same"))
}
