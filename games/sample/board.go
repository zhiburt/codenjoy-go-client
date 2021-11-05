package sample

import (
    "fmt"
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "reflect"
    "sort"
)


type Board struct {
    board *engine.GameBoard
}

func NewBoard(message string) *Board {
    ElementsValues := make([]rune, 0, len(Elements))
    for _, e := range Elements {
        ElementsValues = append(ElementsValues, e)
    }
    return &Board{engine.NewGameBoard(ElementsValues, message)}
}

func (b *Board) IsGameOver() bool {
    return len(b.board.Find(Elements["DEAD_HERO"])) != 0
}


func (b *Board) String() string {
    return b.board.String()
}
