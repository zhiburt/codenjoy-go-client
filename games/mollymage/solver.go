package mollymage

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
	B *Board
}

func (s Solver) Get(rawBoard []rune) string {
	s.B.UpdateBoard(rawBoard)
	fmt.Printf("\nBoard:\n%s\n", s.B.BoardAsString())
	fmt.Printf("Hero at: %v\n", s.B.GetHero())
	fmt.Printf("Other heroes at: %v\n", s.B.GetOtherHeroes())
	fmt.Printf("Ghosts at: %v\n", s.B.GetGhosts())
	fmt.Printf("Treasure Boxes at: %v\n", s.B.GetTreasureBoxes())
	fmt.Printf("Potions as: %v\n", s.B.GetPotions())
	fmt.Printf("Blasts at: %v\n", s.B.GetBlasts())
	fmt.Printf("Expected blasts at: %v\n", s.B.GetFutureBlasts())

	answer := s.nextStep()
	fmt.Println("Answer: " + answer)
	fmt.Println("-------------------------------------------------------------")
	return string(answer)
}

func (s Solver) nextStep() Action {
	// make your action
	return MoveFire(engine.UP)
}
