package engine

import (
	"fmt"
)

type Point struct {
	x, y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) IsValid(boardSize int) bool {
	return (p.x >= 0 && p.x < boardSize) && (p.y >= 0 && p.y < boardSize)
}

func (p *Point) String() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

func StepRight(pt *Point) *Point {
	return NewPoint(pt.X()+1, pt.Y())
}

func StepLeft(pt *Point) *Point {
	return NewPoint(pt.X()-1, pt.Y())
}

func StepUp(pt *Point) *Point {
	return NewPoint(pt.X(), pt.Y()+1)
}

func StepDown(pt *Point) *Point {
	return NewPoint(pt.X(), pt.Y()-1)
}

type SortedPoints []*Point

func (p SortedPoints) Len() int           { return len(p) }
func (p SortedPoints) Less(i, j int) bool { return p[i].String() < p[j].String() }
func (p SortedPoints) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
