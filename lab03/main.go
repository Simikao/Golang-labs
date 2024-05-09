package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State int

// constatnt to decide what to run
// 0 standard mode
// 1 sim mode
const runMode = 0

// state types
const (
	none State = iota
	alive
	burning
	dead
	lightning
)

// Forest variables
const (
	Width   = 20
	Height  = 11
	Density = 0.65
)

// Tree stats
const (
	// types of trees to generate
	speciesSelector = 3

	// deciduous trees 🌳
	DBurning    = 6
	DResistance = 0.55

	// coniferous trees 🌲
	CBurning    = 4
	CResistance = 0.3

	// palm trees 🌴
	PBurning    = 2   // no data
	PResistance = 0.7 // no data
)

type results struct {
	treeSurvived    int
	treeDead        int
	percentSurvived int
	iterations      int
}

func (t results) String() {
	fmt.Sprint("[", t.treeSurvived, t.treeDead, t.percentSurvived, t.iterations, "]")
}

// var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var rng = rand.New(rand.NewSource(1))

var treesAmount int

type Tree struct {
	state          State
	species        int
	maxBurningTime int
	resistance     float64
	burning        int
	position       position
}

func (t Tree) String() string {
	switch t.state {
	case alive:
		switch t.species {
		case 0:
			return "🌳"
		case 1:
			return "🌲"
		case 2:
			return "🌴"
		default:
			return "❓"
		}
	case burning:
		return "🔥"
	case dead:
		return "🪵"
	case lightning:
		return "⚡"
	default:
		return "  "
	}
}

func (t Tree) setTree(state State) Tree {
	t.state = state
	return t
}

func (t Tree) update() Tree {
	if t.state == burning {
		t.burning += 1

		if t.burning >= t.maxBurningTime {
			t.state = dead
		}
	}

	if t.state == lightning {
		t.state = burning
	}
	return t
}

func (t Tree) willIgnite() Tree {
	if rng.Float64() > t.resistance {
		t.state = burning
	}

	return t
}

type position struct {
	x, y int
}

func fromRandomTree(random int) (int, int, float64) {
	switch random {
	case 0:
		return random, DBurning, DResistance
	case 1:
		return random, CBurning, CResistance
	case 2:
		return random, PBurning, PResistance
	default:
		return 0, DBurning, DResistance
	}
}

func generateForest(width int, height int, species int, density float64) [][]Tree {
	forest := make([][]Tree, height)
	for i := range forest {
		forest[i] = make([]Tree, width)
		for j := range forest[i] {
			if rng.Float64() < density {
				speciesType, burningTime, resistance := fromRandomTree(rng.Intn(species))
				forest[i][j] = Tree{alive, speciesType, burningTime, resistance, 0, position{i, j}}
				treesAmount += 1
			} else {
				forest[i][j] = Tree{none, -1, -1, -1, 0, position{i, j}}
			}
		}
	}
	return forest
}
func printForest(forest *[][]Tree) {
	for _, row := range *forest {
		for _, cell := range row {
			fmt.Print(Tree.String(cell))
		}
		fmt.Println()
	}
}

func spreadFire(forest *[][]Tree, i, j int) {
	directions := [8]struct{ di, dj int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		ni, nj := i+dir.di, j+dir.dj

		// Handle potential out-of-bounds access
		if ni < 0 || ni >= Height || nj < 0 || nj >= Width {
			continue
		}

		if (*forest)[ni][nj].state == alive {
			(*forest)[ni][nj] = (*forest)[ni][nj].willIgnite()
		}
	}
}

func burningLoop(forest *[][]Tree) int {
	iterations := 0
	running := true
	for running {
		printForest(forest)
		running = false
		for i, row := range *forest {
			for j, tree := range row {
				(*forest)[i][j] = tree.update()
			}
		}
		for i, row := range *forest {
			for j, _ := range row {
				if (*forest)[i][j].state == burning {
					running = true

					if (*forest)[i][j].burning > 0 {
						spreadFire(forest, i, j)
					}
				}
			}
		}
		iterations += 1
		time.Sleep(500 * time.Millisecond)
	}
	return iterations
}

func thunder(forest *[][]Tree) {
	row := 0
	depth := 0
	for true {
		row = rng.Intn(Height - 1)
		depth = rng.Intn(Width - 1)

		printForest(forest)
		if (*forest)[row][depth].state == alive {
			(*forest)[row][depth].state = lightning
			break
		}
	}
}

func countDead(forest *[][]Tree) int {
	var deadTrees int
	for _, row := range *forest {
		for _, tree := range row {
			if tree.state == dead {
				deadTrees += 1
			}
		}
	}
	return deadTrees
}

func simBurn(width int, height int, species int, density float64) results {
	toAverage := make([]results, 20)
	for x := range 20 {
		forest := generateForest(width, height, species, density)
		fmt.Println("x", x)
		iterations := 0
		running := true
		fmt.Println("Zeus is angy")
		thunder(&forest)
		fmt.Println("boom")
		for running {
			running = false
			for i, row := range forest {
				for j, tree := range row {
					// fmt.Println(i, j)
					forest[i][j] = tree.update()
					// fmt.Println("x", x, "iter:", iterations, "Inm a updateerrrrr r aAAAA")
				}
			}
			fmt.Println("about to spread")
			for i, row := range forest {
				for j := range row {
					if forest[i][j].state == burning {
						running = true

						if forest[i][j].burning > 0 {
							spreadFire(&forest, i, j)
						}
					}
				}
			}
			iterations += 1
			if iterations > 200 {
				fmt.Println("eepy time")
				break
			}
			fmt.Println("inside running", iterations)
		}

		fmt.Println("outside running: ", iterations)
		deadAmount := countDead(&forest)
		fmt.Println("counted")
		toAverage = append(toAverage, results{
			treesAmount,
			deadAmount,
			(100 - (deadAmount * 100 / treesAmount)),
			iterations,
		})
		fmt.Println("Averaged")
	}
	fmt.Println("COllecting results")
	treesAlive := 0
	deadAmount := 0
	percent := 0
	iterations := 0
	for _, elem := range toAverage {
		treesAlive += elem.treeSurvived / 20
		deadAmount += elem.treeDead / 20
		percent += elem.percentSurvived / 20
		iterations += elem.iterations / 20
	}

	fmt.Println("PRinting results")

	return results{
		treesAlive,
		deadAmount,
		percent,
		iterations,
	}
}
func main() {
	forest := generateForest(Width, Height, speciesSelector, Density)
	printForest(&forest)

	// forest[Height-1][Width-1].state = burning
	// forest[0][0].state = lightning
	switch runMode {
	case 0:

		thunder(&forest)
		iterations := burningLoop(&forest)
		deadAmount := countDead(&forest)
		printForest(&forest)

		fmt.Printf("For %d trees at the beginning %d have burned down. Meaning only %d%% survied\n", treesAmount, deadAmount, (100 - (deadAmount * 100 / treesAmount)))
		fmt.Printf("It took %d iterations\n", iterations)
	case 1:
		output2species := make([]results, 0)
		output3species := make([]results, 0)
		for i := range 20 {
			fmt.Println("Starting")
			output2species = append(output2species, simBurn(20, 20, 2, float64(i+1)*0.05))
			fmt.Println("Am I here")
			output3species = append(output3species, simBurn(20, 20, 3, float64(i+1)*0.05))
			fmt.Println("Bm I here")
			fmt.Println(i)
		}
		fmt.Println("WYSZEDŁEM!")

		fmt.Println("Results for just two species: ")
		for i, elem := range output2species {
			if i%5 == 0 {
				fmt.Println()
			}
			fmt.Print(elem)
		}

		fmt.Println()
		fmt.Println("Results after adding palm trees: ")
		for i, elem := range output3species {
			if i%5 == 0 {
				fmt.Println()
			}
			fmt.Print(elem)
		}
		fmt.Println()
	}

}
