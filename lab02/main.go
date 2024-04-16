package main

import (
	"fmt"
	"math/big"
	"strings"
)

var bigIntOne = big.NewInt(int64(1))

type Person struct {
	Name     string
	LastName string
}

type Factorial struct {
	iteration int64
	value     *big.Int
}

func (f *Factorial) next() {
	f.iteration++
	f.value = f.value.Mul(f.value, big.NewInt(f.iteration))
}

func NewFactorial() *Factorial {
	tmp := new(Factorial)
	tmp.iteration = 0
	tmp.value = big.NewInt(int64(1))
	return tmp
}

func pseudoFactorial(number *big.Int) *big.Int {
	return number.Mul(number.Add(number, bigIntOne), number)
}

func isStrongNumber(bigNumber *Factorial, nick string) bool {
	found := true
	for _, char := range nick {
		codeFinal := fmt.Sprint(char)
		if !strings.Contains(bigNumber.value.String(), codeFinal) {
			found = false
		}
	}

	return found
}

func findStrongNumber(nick string) {
	number := NewFactorial()
	for !isStrongNumber(number, nick) {
		number.next()
	}

}

func main() {
	me := Person{"John", "Doe"}
	nick := string(strings.ToLower(me.Name[0:3]))
	nick += string(strings.ToLower(me.LastName[0:3]))
	fmt.Println(nick)
	test := NewFactorial()
	fmt.Println(bigIntOne)
	fmt.Println(test.iteration, test.value)
	for !isStrongNumber(test, nick) {
		test.next()
	}
	fmt.Println(test.iteration, test.value)

	test2 := NewFactorial()
	for !isStrongNumber(test2, "pioarl") {
		test2.next()
	}
	fmt.Println(test2.iteration, test.value)
}
