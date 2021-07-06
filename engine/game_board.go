package engine

import (
	"fmt"
	"math"
	"strings"
)

type GameBoard struct {
	elements []rune
	len      int
	size     int
}

func NewGameBoard(supportedElements []rune, message string) *GameBoard {
	board := GameBoard{}
	message = strings.Replace(message, "board=", "", 1)
	board.initElementsArray(supportedElements, []rune(message))
	board.len = len(board.elements)
	board.size = int(math.Sqrt(float64(board.len)))
	return &board
}

func (b *GameBoard) initElementsArray(supportedElements []rune, rawBoard []rune) {
	b.elements = make([]rune, len(rawBoard))
	for i := 0; i < len(b.elements); i++ {
		nextElement := rawBoard[i]
		for _, v := range supportedElements {
			if nextElement == v {
				b.elements[i] = nextElement
				break
			}
		}
		if b.elements[i] == 0 {
			panic(fmt.Sprintf("invalid element: %v", nextElement))
		}
	}
}

func (b *GameBoard) GetSize() int {
	return b.size
}

func (b *GameBoard) GetAt(pt *Point) rune {
	if !pt.IsValid(b.size) {
		panic(fmt.Sprintf("invalid point %s" + pt.String()))
	}
	return b.elements[b.pointToIndex(pt.X(), pt.Y())]
}

func (b *GameBoard) Find(wanted ...rune) []*Point {
	var points []*Point
	for _, w := range wanted {
		for i, el := range b.elements {
			if w == el {
				points = append(points, b.indexToPoint(i))
			}
		}
	}
	return points
}

func (b *GameBoard) FindFirst(wanted ...rune) *Point {
	for _, w := range wanted {
		for j, el := range b.elements {
			if w == el {
				return b.indexToPoint(j)
			}
		}
	}
	return nil
}

func (b *GameBoard) IsAt(pt *Point, wanted ...rune) bool {
	if !pt.IsValid(b.size) {
		return false
	}
	el := b.GetAt(pt)
	for _, w := range wanted {
		if w == el {
			return true
		}
	}
	return false
}

func (b *GameBoard) FindNear(pt *Point) []rune {
	var elements []rune

	right := NewPoint(pt.X()+1, pt.Y())
	if right.IsValid(b.size) {
		elements = append(elements, b.GetAt(right))
	}
	left := NewPoint(pt.X()-1, pt.Y())
	if left.IsValid(b.size) {
		elements = append(elements, b.GetAt(left))
	}
	up := NewPoint(pt.X(), pt.Y()+1)
	if up.IsValid(b.size) {
		elements = append(elements, b.GetAt(up))
	}
	down := NewPoint(pt.X()+1, pt.Y()-1)
	if down.IsValid(b.size) {
		elements = append(elements, b.GetAt(down))
	}

	return elements
}

func (b *GameBoard) CountNear(pt *Point, wanted ...rune) int {
	counter := 0
	for _, el := range b.FindNear(pt) {
		for _, w := range wanted {
			if w == el {
				counter++
			}
		}
	}
	return counter
}

func (b *GameBoard) IsNear(pt *Point, wanted ...rune) bool {
	for _, w := range wanted {
		if b.CountNear(pt, w) != 0 {
			return true
		}
	}
	return false
}

func (b *GameBoard) pointToIndex(x int, y int) int {
	return (b.size-1-y)*b.size + x
}

func (b *GameBoard) indexToPoint(index int) *Point {
	x := index % b.size
	y := int(math.Ceil(float64(b.size - 1 - index/b.size)))
	return NewPoint(x, y)
}

func (b *GameBoard) String() string {
	builder := strings.Builder{}
	for y := b.size - 1; y >= 0; y-- {
		for x := 0; x < b.size; x++ {
			builder.WriteString(string(b.elements[b.pointToIndex(x, y)]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
