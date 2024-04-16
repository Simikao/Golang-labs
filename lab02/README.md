# Wprowadzenie

To zadanie polega na obliczeniu dwóch powiązanych liczb, które można wygenerować na podstawie Twojego imienia i nazwiska. Wymaga ono wykonania dużej ilości obliczeń na bardzo dużych liczbach.

# Część 1: Generowanie Nicku
```go

package main

import "strings"

var (
    replacer = strings.NewReplacer(
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

func main() {
    // ... (Kod wprowadzania danych)

    nick := string(strings.ToLower(replacer.Replace(user.Name)[0:3]))
    nick += string(strings.ToLower(replacer.Replace(user.LastName)[0:3]))

    // ... (Reszta kodu)
}
```

## Wyjaśnienie:
1. Importujemy pakiet `strings`, który będzie używany do manipulacji łańcuchami znaków.
2. Definiujemy zmienną `replacer` typu `*strings.Replacer`, która zawiera słownik mapowania polskich znaków diakrytycznych na ich odpowiedniki ASCII.
3. W funkcji `main` pobieramy dane użytkownika (imię i nazwisko).
4. Używamy funkcji `strings.ToLower` i `replacer.Replace` do konwersji nicku na małe litery i zamiany polskich znaków na ASCII.
3. Łączymy pierwsze 3 litery imienia i nazwiska, aby utworzyć `nick`.

# Część 2: Silna Liczba

```go
Go

package main

import (
    "big/int"
    "fmt"
)

type Factorial struct {
    iteration int64
    value      *big.Int
}

func NewFactorial() *Factorial {
    tmp := new(Factorial)
    tmp.iteration = 0
    tmp.value = big.NewInt(int64(1))
    return tmp
}

func (f *Factorial) next() {
    f.iteration++
    f.value = f.value.Mul(f.value, big.NewInt(f.iteration))
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
```
## Wyjaśnienie:
1. Importujemy pakiety `big/int` i `fmt` do obsługi dużych liczb i formatowania tekstu.
1. Definiujemy strukturę `Factorial` do przechowywania informacji o wartości silni i liczbie iteracji.
1. Funkcja `NewFactorial` tworzy nowy obiekt `Factorial` z początkową wartością 1.
1. Metoda `next` struktury `Factorial` oblicza kolejną wartość silni mnożąc bieżącą wartość przez liczbę iteracji.
1. Funkcja `isStrongNumber` sprawdza, czy dana wartość silni zawiera wszystkie kody ASCII z nicku.
1. Funkcja `findStrongNumber` iteruje, obliczając kolejne wartości silni, aż znajdzie taką, która spełnia warunek `isStrongNumber`.

# Część 3: Słaba Liczba

```go
package main

func fibonacci(n int, arr *[]int) int {
    (*arr)[n] += 1
    if n <= 1 {
        return n
    }
    return fibonacci(n-1, arr) + fibonacci(n-2, arr)
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
    // ... (Reszta kodu)
    
	arr := make([]int, 31)
	fibonacci(30, &arr)

	weakNumber := findWeakNumber(result.iteration, &arr)
	
    // ... (Wyświetlanie odpowiedzi)
}
```
## Wyjaśnienie:
1. Funkcja `fibonacci` oblicza wartość n-tego elementu ciągu Fibonacciego.
1. Używa tablicy `arr` do zliczania liczby wywołań funkcji dla każdego argumentu.
1. Funkcja `findWeakNumber` szuka argumentu m w ciągu Fibonacciego, którego liczba wywołań jest najbliższa wartości `strongNumber`.
1. Konwertuje `strongNumber` na typ int i oblicza początkową różnicę.
1. Przechodzi przez tablicę `arr` od końca do początku.
1. Oblicza nową różnicę dla każdego argumentu m.
1. Jeśli nowa różnica jest mniejsza lub równa bieżącej różnicy, aktualizuje `difference` i `weakNumber`.
1. Jeśli nowa różnica jest większa, zwraca aktualny `weakNumber`.
1. Jeśli nie znajdzie odpowiedniego argumentu, zwraca `-1`.



# Przykładowy output dla znanej już wartości z zdania

![Example](https://i.ibb.co/rZ17Ht4/image.png "example output")
