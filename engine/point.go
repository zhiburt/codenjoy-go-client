package engine

import "fmt"

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
	return (p.x >= 0 && p.x <= boardSize) && (p.y >= 0 && p.y <= boardSize)
}

func (p Point) String() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}
