package main

import (
	"bytes"
	"fmt"

	"github.com/dlokkers/gophercises/phone/phoneStore"
)

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func main() {

	numbers := []string{"1234567890","123 456 7891","(123) 456 7892","(123) 456-7893","123-456-7894","123456789","(123)456-7895"}
	for _, i := range numbers {
		_, err := phones.InsertPhone(i)
		if err != nil {
			panic(err)
		}
	}

	retrievedNumbers, err := phones.GetAllNumbers()
	if err != nil {
		panic(err)
	}

	for _, i := range retrievedNumbers {
		normalized := normalize(i.Number)
		fmt.Printf("%d - %s : %s\n", i.Id, i.Number, normalized)
	}
	// normalize numbers
	// store numbers
}
