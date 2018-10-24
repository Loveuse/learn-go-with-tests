package iteration

import "fmt"

const repeatLimit = 5

func Hello() string {
	return "Hello, World!"
}

func Repeat(character string) (repeatString string) {
	for i := 0; i < repeatLimit; i++ {
		repeatString += character
	}
	return
}

func main() {
	fmt.Println(Hello())
}
