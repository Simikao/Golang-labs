package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

var (
	debugMode     = false
	fibbonaciTime = false
	ackermannDo   = false
	bigIntOne     = big.NewInt(int64(1))
	replacer      = strings.NewReplacer(
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

func timeFunction(number int, function func(int, *[]int) int) string {
	slice := make([]int, number+1)
	startTime := time.Now()
	function(number, &slice)
	stopTime := time.Now()

	return stopTime.Sub(startTime).String()
}

func getName() (string, string) {

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

	return name, lastName

}

func ackermann(m, n int) int {

	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}

func main() {

	var name string
	var lastName string

	if !debugMode {
		name, lastName = getName()
	} else {

		name = "Piorun"
		lastName = "Arłeński"
	}
	user := Person{name, lastName}

	nick := string(strings.ToLower(replacer.Replace(user.Name)[0:3]))
	nick += string(strings.ToLower(replacer.Replace(user.LastName)[0:3]))

	result := findStrongNumber(nick)
	fmt.Printf("For the name %s %s, whose nick is %q, the strong number is %d\n", user.Name, user.LastName, nick, result.iteration)

	arr := make([]int, 31)
	fibonacci(30, &arr)

	weakNumber := findWeakNumber(result.iteration, &arr)
	if weakNumber < 0 {
		fmt.Println("Something went wrong")
	} else {
		fmt.Printf("Weak number for %s %s is %d and it was called %d times in the recursive fibonacci function (from number 30)\n", user.Name, user.LastName, weakNumber, arr[weakNumber])
	}

	if fibbonaciTime {
		fmt.Println("10", timeFunction(10, fibonacci))
		fmt.Println("15", timeFunction(15, fibonacci))
		fmt.Println("20", timeFunction(20, fibonacci))
		fmt.Println("25", timeFunction(25, fibonacci))
		fmt.Println("30", timeFunction(30, fibonacci))
		fmt.Println("35", timeFunction(35, fibonacci))
		fmt.Println("40", timeFunction(40, fibonacci))
		fmt.Println("45", timeFunction(45, fibonacci))
		fmt.Println("50", timeFunction(50, fibonacci))
		fmt.Println("55", timeFunction(55, fibonacci))
	}

	if ackermannDo {

		startTime := time.Now()
		num := ackermann(0, 0)
		stopTime := time.Now()
		end := stopTime.Sub(startTime).String()
		fmt.Println("ack 0,0 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(0, 1)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 0,1 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(1, 0)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 1,0 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(1, 1)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 1,1 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(1, 2)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 1,2 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(2, 2)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 2,2 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(3, 2)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 3,2 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(4, 0)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 4,0 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(4, 1)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 4,1 took ", end, " time")
		fmt.Println("result: ", num)

		startTime = time.Now()
		num = ackermann(4, 2)
		stopTime = time.Now()
		end = stopTime.Sub(startTime).String()
		fmt.Println("ack 4,2 took ", end, " time")
		fmt.Println("result: ", num)
	}

}
