package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
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
		"\n", "",
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
func fibonacci(n int, arr *[]int) int {
	(*arr)[n] += 1
	if n <= 1 {
		return n
	}
	return fibonacci(n-1, arr) + fibonacci(n-2, arr)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findWeakNumber(strongNumber int64, arr *[]int) int {
	intSN := int(strongNumber)
	difference := abs(intSN - (*arr)[len(*arr)-1])
	for i := (len(*arr) - 1); i >= 0; i-- {
		newDifference := abs(intSN - (*arr)[i])
		if newDifference <= difference {
			difference = newDifference
		} else {
			return i + 1
		}
	}
	return -1
}
func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong, check your input and try again")
	}
	name = strings.Replace(name, "\n", "", -1)

	fmt.Print("Enter last name: ")
	lastName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong, check your input and try again")
	}
	lastName = strings.Replace(lastName, "\n", "", -1)

	user := Person{name, lastName}

	nick := string(strings.ToLower(replacer.Replace(user.Name)[0:3]))
	nick += string(strings.ToLower(replacer.Replace(user.LastName)[0:3]))

	result := findStrongNumber(nick)
	fmt.Printf("For the name %s %s, whose nick is %q, the strong number is %d\n", user.Name, user.LastName, nick, result.iteration)

	arr := make([]int, 31)
	fibonacci(30, &arr)

	weakNumberR := findWeakNumber(result.iteration, &arr)
	weakNumber := len(arr) - weakNumberR
	if weakNumberR < 0 {
		fmt.Println("Something went wrong")
	} else {
		fmt.Printf("Weak number for %s %s is %d and it was called %d times in the recursive fibonacci function (from number 30)\n", user.Name, user.LastName, weakNumber, arr[weakNumberR])
	}

}
