package clifford

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
    fmt.Println("\nAnswer: " + action.String())
    fmt.Println("-------------------------------------------------------------")
    return action.String()
}

func (s *Solver) nextAction(b *Board) engine.Direction {
    // TODO: write your code here
    return engine.ACT
}
