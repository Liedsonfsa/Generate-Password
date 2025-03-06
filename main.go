package main

import (
	"fmt"

	"github.com/Liedsonfsa/Generate-Password/password"
)

func main() {
	generatedPassword := password.Generate()

	fmt.Println(generatedPassword)
}