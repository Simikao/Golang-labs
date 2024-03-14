package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	winning = 1 << iota
	chosen
	revealed
)

func chooseFirst(board *[3]int) bool {
	selected := rand.Intn(3)
	(*board)[selected] += chosen
	return (*board)[selected] == 3
}

func indexSelected(s *[3]int) int {
	for i, vs := range s {
		if vs >= chosen && vs < revealed {
			return i
		}
	}

	return -1
}

func findZeroIndices(arr [3]int) []int {
	var zeroIndices []int

	for i, v := range arr {
		if v == 0 {
			zeroIndices = append(zeroIndices, i)
		}
	}

	return zeroIndices
}

func theGame(strategy bool) int {
	goal := rand.Intn(3)
	board := [...]int{0, 0, 0}
	board[goal] = winning
	// fmt.Println(board)

	chooseFirst(&board)
	// fmt.Println(board)

	selectedByUser := indexSelected(&board)
	if selectedByUser < 0 {
		fmt.Println("Something went horribly wrong")
		return -10000000
	}

	if strategy {
		toReveal := findZeroIndices(board)
		revealedIndex := toReveal[rand.Intn(len(toReveal))]
		// fmt.Println("Randomly selected index:", revealedIndex)
		board[revealedIndex] += revealed

		// fmt.Println("selected", selectedByUser)
		leftIndex := 0 + 1 + 2 - (revealedIndex + selectedByUser)
		// fmt.Println(leftIndex)

		board[selectedByUser] -= chosen
		board[leftIndex] += chosen

		selectedByUser = leftIndex
	}

	// fmt.Println("Final board:", board, "user chose", selectedByUser, "which is", board[selectedByUser])
	if board[selectedByUser] == 3 {
		// fmt.Println("you won!")
		return 1
	} else {
		// fmt.Println("You lost D:")
		return 0
	}
}

func main() {
	rounds := 1000
	withRounds := rounds / 2
	withoutRound := rounds - withRounds

	var wg sync.WaitGroup

	changeWon := 0
	noChangeWon := 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range withRounds {
			changeWon += theGame(true)
		}
	}()

	go func() {
		defer wg.Done()
		for range withoutRound {
			noChangeWon += theGame(false)
		}
	}()

	wg.Wait()
	fmt.Printf("Na %d prób takie są wyniki:\n", rounds)
	fmt.Printf("Jeśli osoba zmieniła walizkę to wygrała %d razy a przegrała %d razy\n", changeWon, withRounds-changeWon)
	fmt.Printf("Jeśli osoba nie zmieniła walizki to wygrała %d razy a przegrała %d razy\n", noChangeWon, withRounds-noChangeWon)
}
