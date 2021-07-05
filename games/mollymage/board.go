package mollymage

import (
	"errors"
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

const BLAST_SIZE int = 3

type Board struct {
	board *engine.GameBoard
}

func NewBoard(message string) *Board {
	elementsValues := make([]string, 0, len(elements))
	for _, tx := range elements {
		elementsValues = append(elementsValues, tx)
	}
	return &Board{engine.NewGameBoard(elementsValues, message)}
}

func (b *Board) GetAt(pt *engine.Point) string {
	if !pt.IsValid(b.board.GetSize()) {
		return elements["WALL"]
	}
	el, _ := b.board.GetAt(pt)
	return el
}

func (b *Board) FindHero() (*engine.Point, error) {
	points := b.board.Find(elements["HERO"],
		elements["POTION_HERO"],
		elements["DEAD_HERO"])
	if len(points) == 0 {
		return &engine.Point{}, errors.New("hero element has not been found")
	}
	return points[0], nil
}

func (b *Board) IsGameOver() bool {
	return len(b.board.Find(elements["DEAD_HERO"])) != 0
}

func (b *Board) FindOtherHeroes() []*engine.Point {
	return b.board.Find(elements["OTHER_HERO"],
		elements["OTHER_POTION_HERO"],
		elements["OTHER_DEAD_HERO"])
}

func (b *Board) FindBarriers() []*engine.Point {
	var points []*engine.Point
	points = append(points, b.FindWalls()...)
	points = append(points, b.FindGhosts()...)
	points = append(points, b.FindTreasureBoxes()...)
	points = append(points, b.FindPotions()...)
	points = append(points, b.FindOtherHeroes()...)
	return points
}

func (b *Board) FindWalls() []*engine.Point {
	return b.board.Find(elements["WALL"])
}

func (b *Board) FindGhosts() []*engine.Point {
	return b.board.Find(elements["GHOST"])
}

func (b *Board) FindTreasureBoxes() []*engine.Point {
	return b.board.Find(elements["TREASURE_BOX"])
}

func (b *Board) FindPotions() []*engine.Point {
	return b.board.Find(elements["POTION_TIMER_1"],
		elements["POTION_TIMER_2"],
		elements["POTION_TIMER_3"],
		elements["POTION_TIMER_4"],
		elements["POTION_TIMER_5"],
		elements["POTION_HERO"],
		elements["OTHER_POTION_HERO"])
}

func (b *Board) FindBlasts() []*engine.Point {
	return b.board.Find(elements["BOOM"])
}

func (b *Board) PredictFutureBlasts() []*engine.Point {
	// TODO: implement
	return []*engine.Point{}
}

func (b *Board) FindPerks() []*engine.Point {
	return b.board.Find(elements["POTION_COUNT_INCREASE"],
		elements["POTION_REMOTE_CONTROL"],
		elements["POTION_IMMUNE"],
		elements["POTION_BLAST_RADIUS_INCREASE"])
}

func (b *Board) String() string {
	hero, _ := b.FindHero()
	return b.board.String() +
		"\nHero at: " + hero.String() +
		"\nOther heroes at: " + fmt.Sprintf("%v", b.FindOtherHeroes()) +
		"\nGhosts at: " + fmt.Sprintf("%v", b.FindGhosts()) +
		"\nPotions at: " + fmt.Sprintf("%v", b.FindPotions()) +
		"\nBlasts at: " + fmt.Sprintf("%v", b.FindBlasts()) +
		"\nExpected blasts at: " + fmt.Sprintf("%v", b.PredictFutureBlasts())
}
