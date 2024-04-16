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

const (
	rounds = 1
	doors  = 200
)

type gameBoard [doors]int

func chooseFirst(board *gameBoard) {
	selected := rand.Intn(doors)
	(*board)[selected] += chosen
}

func indexSelected(s *gameBoard) int {
	for i, vs := range s {
		if vs >= chosen && vs < revealed {
			return i
		}
	}

	return -1
}

func findZeroIndices(arr *gameBoard) []int {
	var zeroIndices []int

	for i, v := range arr {
		if v == 0 {
			zeroIndices = append(zeroIndices, i)
		}
	}

	return zeroIndices
}

func findNotRevealedIndices(arr *gameBoard) []int {
	var opened []int
	for i, v := range arr {
		if v != revealed {
			opened = append(opened, i)
		}
	}
	return opened
}

func theGame(strategy bool) int {
	goal := rand.Intn(doors)
	board := gameBoard{}
	board[goal] = winning
	// fmt.Println(board)

	chooseFirst(&board)
	fmt.Println(board)

	selectedByUser := indexSelected(&board)
	if selectedByUser < 0 {
		fmt.Println("Something went horribly wrong")
		return -10000000
	}

	if strategy {
		toReveal := findZeroIndices(&board)
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

func theOtherGame(revealedDoors int) int {
	goal := rand.Intn(doors)
	board := gameBoard{}
	board[goal] = winning
	// fmt.Println(board)

	chooseFirst(&board)
	// fmt.Println(board)

	selectedByUser := indexSelected(&board)
	if selectedByUser < 0 {
		fmt.Println("Something went horribly wrong")
		return -10000000
	}

	toReveal := findZeroIndices(&board)
	// fmt.Println("Randomly selected index:", revealedIndex)
	for i := 0; i < revealedDoors; i++ {
		revealedIndex := toReveal[rand.Intn(len(toReveal))]
		if board[revealedIndex] == 0 {
			board[revealedIndex] += revealed
		} else {
			i--
		}

	}

	toChange := findNotRevealedIndices(&board)
	for i := 0; i < len(board); i++ {
		toChangeIndex := toChange[rand.Intn(len(toChange))]
		if toChangeIndex != selectedByUser {
			board[toChangeIndex] += chosen
			board[selectedByUser] -= chosen
			selectedByUser = toChangeIndex
			break
		}
	}

	if board[selectedByUser] == 3 {
		return 1
	} else {
		return 0
	}

	// fmt.Println("selected", selectedByUser)
	// leftIndex := 0 + 1 + 2 - (revealedIndex + selectedByUser)
	// fmt.Println(leftIndex)

	// board[selectedByUser] -= chosen
	// board[leftIndex] += chosen

	// selectedByUser = leftIndex

}

func main() {
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

	results := make(map[int]int)
	for revealing := range doors - 2 {
		gamesWon := 0
		func() {
			for range rounds {
				gamesWon += theOtherGame(revealing + 1)
			}
			results[revealing+1] = gamesWon
		}()

	}
	wg.Wait()
	fmt.Printf("Na %d prób takie są wyniki:\n", rounds)
	fmt.Printf("Jeśli osoba zmieniła walizkę to wygrała %d razy a przegrała %d razy\n", changeWon, withRounds-changeWon)
	fmt.Printf("Jeśli osoba nie zmieniła walizki to wygrała %d razy a przegrała %d razy\n", noChangeWon, withRounds-noChangeWon)
	fmt.Println(results)
}
