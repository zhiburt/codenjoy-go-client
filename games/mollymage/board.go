package mollymage

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"reflect"
	"sort"
)

const BLAST_SIZE int = 3

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

func (b *Board) GetAt(pt *engine.Point) rune {
	if !pt.IsValid(b.board.GetSize()) {
		return Elements["WALL"]
	}
	return b.board.GetAt(pt)
}

func (b *Board) FindHero() *engine.Point {
	points := b.board.Find(Elements["HERO"],
		Elements["POTION_HERO"],
		Elements["DEAD_HERO"])
	if len(points) == 0 {
		panic("hero element has not been found")
	}
	return points[0]
}

func (b *Board) IsGameOver() bool {
	return len(b.board.Find(Elements["DEAD_HERO"])) != 0
}

func (b *Board) FindOtherHeroes() []*engine.Point {
	return b.board.Find(Elements["OTHER_HERO"],
		Elements["OTHER_POTION_HERO"],
		Elements["OTHER_DEAD_HERO"])
}

func (b *Board) FindBarriers() []*engine.Point {
	var points []*engine.Point
	points = appendIfMissing(points, b.FindWalls()...)
	points = appendIfMissing(points, b.FindGhosts()...)
	points = appendIfMissing(points, b.FindTreasureBoxes()...)
	points = appendIfMissing(points, b.FindPotions()...)
	points = appendIfMissing(points, b.FindOtherHeroes()...)
	sort.Sort(engine.SortedPoints(points))
	return points
}

func appendIfMissing(slice []*engine.Point, points ...*engine.Point) []*engine.Point {
	for _, p := range points {
		existed := false
		for _, ele := range slice {
			if reflect.DeepEqual(ele, p) {
				existed = true
				break
			}
		}
		if !existed {
			slice = append(slice, p)
		}
	}
	return slice
}

func (b *Board) FindWalls() []*engine.Point {
	return b.board.Find(Elements["WALL"])
}

func (b *Board) FindGhosts() []*engine.Point {
	return b.board.Find(Elements["GHOST"])
}

func (b *Board) FindTreasureBoxes() []*engine.Point {
	return b.board.Find(Elements["TREASURE_BOX"])
}

func (b *Board) FindPotions() []*engine.Point {
	return b.board.Find(Elements["POTION_TIMER_1"],
		Elements["POTION_TIMER_2"],
		Elements["POTION_TIMER_3"],
		Elements["POTION_TIMER_4"],
		Elements["POTION_TIMER_5"],
		Elements["POTION_HERO"],
		Elements["OTHER_POTION_HERO"])
}

func (b *Board) FindBlasts() []*engine.Point {
	return b.board.Find(Elements["BOOM"])
}

func (b *Board) PredictFutureBlasts() []*engine.Point {
	// TODO: implement
	return []*engine.Point{}
}

func (b *Board) FindPerks() []*engine.Point {
	return b.board.Find(Elements["POTION_COUNT_INCREASE"],
		Elements["POTION_REMOTE_CONTROL"],
		Elements["POTION_IMMUNE"],
		Elements["POTION_BLAST_RADIUS_INCREASE"])
}

func (b *Board) String() string {
	return b.board.String() +
		"\nHero at: " + b.FindHero().String() +
		"\nOther heroes at: " + fmt.Sprintf("%v", b.FindOtherHeroes()) +
		"\nGhosts at: " + fmt.Sprintf("%v", b.FindGhosts()) +
		"\nPotions at: " + fmt.Sprintf("%v", b.FindPotions()) +
		"\nBlasts at: " + fmt.Sprintf("%v", b.FindBlasts()) +
		"\nExpected blasts at: " + fmt.Sprintf("%v", b.PredictFutureBlasts())
}
