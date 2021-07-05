package engine

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type GameBoard struct {
	elements []string
	len      int
	size     int
}

func NewGameBoard(supportedElements []string, message string) *GameBoard {
	board := GameBoard{}
	message = strings.Replace(message, "board=", "", 1)
	board.initElementsArray(supportedElements, message)
	board.len = len(board.elements)
	board.size = int(math.Sqrt(float64(board.len)))
	return &board
}

func (b *GameBoard) initElementsArray(supportedElements []string, message string) error {
	b.elements = make([]string, len(message))
	for i := 0; i < len(b.elements); i++ {
		nextElement := string(message[i])
		for _, v := range supportedElements {
			if nextElement == v {
				b.elements[i] = nextElement
				break
			} else {
				return errors.New(fmt.Sprintf("invalid element: %s", nextElement))
			}
		}
	}
	return nil
}

func (b *GameBoard) GetSize() int {
	return b.size
}

func (b *GameBoard) GetAt(pt *Point) (string, error) {
	if !pt.IsValid(b.size) {
		return "", errors.New(fmt.Sprintf("invalid point %s" + pt.String()))
	}
	return b.elements[b.pointToIndex(pt.X(), pt.Y())], nil
}

func (b *GameBoard) Find(wanted ...string) []*Point {
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

func (b *GameBoard) FindFirst(wanted ...string) *Point {
	for _, w := range wanted {
		for j, el := range b.elements {
			if w == el {
				return b.indexToPoint(j)
			}
		}
	}
	return nil
}

func (b *GameBoard) IsAt(pt *Point, wanted ...string) bool {
	if !pt.IsValid(b.size) {
		return false
	}
	el, err := b.GetAt(pt)
	if err != nil {
		return false
	}
	for _, w := range wanted {
		if w == el {
			return true
		}
	}
	return false
}

func (b *GameBoard) FindNear(pt *Point) []string {
	var elements []string

	right := NewPoint(pt.X()+1, pt.Y())
	if right.IsValid(b.size) {
		el, _ := b.GetAt(right)
		elements = append(elements, el)
	}
	left := NewPoint(pt.X()-1, pt.Y())
	if left.IsValid(b.size) {
		el, _ := b.GetAt(left)
		elements = append(elements, el)
	}
	up := NewPoint(pt.X(), pt.Y()+1)
	if up.IsValid(b.size) {
		el, _ := b.GetAt(up)
		elements = append(elements, el)
	}
	down := NewPoint(pt.X()+1, pt.Y()-1)
	if down.IsValid(b.size) {
		el, _ := b.GetAt(down)
		elements = append(elements, el)
	}

	return elements
}

func (b *GameBoard) CountNear(pt *Point, wanted ...string) int {
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

func (b *GameBoard) IsNear(pt *Point, wanted ...string) bool {
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
	}
	return builder.String()
}
