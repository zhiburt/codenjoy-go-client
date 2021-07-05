package mollymage

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
}

func NewSolver() engine.Solver {
	return &Solver{}
}

func (s *Solver) Answer(message string) string {
	board := NewBoard(message)
	fmt.Println("Board \n" + board.String())
	action := s.nextAction(board)
	fmt.Println("\nAnswer: " + action)
	fmt.Println("-------------------------------------------------------------")
	return action
}

func (s *Solver) nextAction(b *Board) string {
	// TODO: write your code here
	return ACT
}
