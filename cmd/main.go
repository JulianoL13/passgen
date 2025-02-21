package main

import (
	keygen "github.com/JulianoL13/passgen/internal/random"
)

func main() {
	test := keygen.GetRandomPass(15, true, true, true)
	println(test)
	println(len(test))
}
