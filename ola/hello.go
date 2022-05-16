package main

import "fmt"

const espanhol = "espanol"
const french = "french"
const eng = "eng"
const prefixPortOi = "Oi "
const prefixEspanhol = "Hola, "
const prefixFrench = "Bonjour, "
const prefixEng = "Hi, "

func Hello(name, idioma string) string {
	if name == "" {
		name = "Any"
	}

	return prefixSaudacao(idioma) + name
}

func prefixSaudacao(idioma string) (prefix string) {
	switch idioma {
	case french:
		prefix = prefixFrench
	case espanhol:
		prefix = prefixEspanhol
	case eng:
		prefix = prefixEng
	}
	return

}

func main() {
	fmt.Println(Hello("anyone", "same"))
}
