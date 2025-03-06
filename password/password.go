package password

import "math/rand"

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

func upperCase() byte {
	var upper string = "ZXCVBNMLSKDJFHGPOQWIEURYT"
	index := rand.Intn(len(upper))

	return byte(upper[index]) 
}

func lowerCase() byte {
	var lower string = "aqzwsxcderfvbgtyhnmjuiklop"
	index := rand.Intn(len(lower))
	
	return byte(lower[index])
}

func number() byte {
	var numbers string = "0123456789"
	index := rand.Intn(len(numbers))

	return byte(numbers[index])
}

func especial() byte {
	var especials string = "!@#$%^&*()_+{}[]:;<>,.?/~`"
	index := rand.Intn(len(especials))

	return byte(especials[index])
}