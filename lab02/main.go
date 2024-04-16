package main

import (
	"fmt"
	"math/big"
	"strings"
)

var (
	bigIntOne = big.NewInt(int64(1))
	replacer  = strings.NewReplacer(
		"ą", "a",
		"ć", "c",
		"ę", "e",
		"ł", "l",
		"ń", "n",
		"ó", "o",
		"ś", "s",
		"ź", "z",
		"ż", "z",
	)
)

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

func findStrongNumber(nick string) *Factorial {
	number := NewFactorial()
	for !isStrongNumber(number, nick) {
		number.next()
	}

	return number
}

func main() {
	me := Person{"Pioterl", "Arłenski"}
	nick := string(strings.ToLower(replacer.Replace(me.Name)[0:3]))
	nick += string(strings.ToLower(replacer.Replace(me.LastName)[0:3]))
	fmt.Println(nick)
	result := findStrongNumber(nick)
	fmt.Printf("For the name %s %s, whose nick is %q, the strong number is %d\n", me.Name, me.LastName, nick, result.iteration)
}
