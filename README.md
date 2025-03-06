# Generate Password

A password generator package

## Package functions

```Go
// Generate generate a password
func Generate() string {
    password := []byte{}
	var char byte

	for range 16 {
		drawnFunction := rand.Intn(4) + 1
		switch drawnFunction {
		case 1:
			char = upperCase()
		case 2:
			char = lowerCase()
		case 3:
			char = number()
		case 4:
			char = especial()
		}

		password = append(password, char)
	}

	return string(password)
}
```

## Downloading package

```bash
go get -u github.com/Liedsonfsa/Generate-Password@v0.0.1
```

## Importing the package

```Go
import "github.com/Liedsonfsa/Generate-Password/password"
```

## Generating the password

```Go
package main

import (
	"github.com/Liedsonfsa/Generate-Password/password"
)

func main() {
	generatedPassword := password.Generate()
}
```