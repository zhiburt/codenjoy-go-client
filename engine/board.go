package engine

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type AbstractBoard struct {
	Content []rune
	Size    int
}

type Element rune

func (b *AbstractBoard) UpdateBoard(content []rune) {
	b.Content = content
	b.Size = int(math.Sqrt(float64(len(content))))
}

func (b *AbstractBoard) BoardAsString() string {
	if len(b.Content) == 0 {
		return ""
	}

	repr := strings.Builder{}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			c := b.Content[i*b.Size+j]
			repr.Write([]byte(fmt.Sprintf("%c", c)))
		}
		repr.Write([]byte("\n"))
	}
	return repr.String()
}

func (b *AbstractBoard) FindOne(symbol Element) (Point, error) {
	for i, el := range b.Content {
		if el == rune(symbol) {
			return b.IndexToPoint(i), nil
		}
	}
	return Point{-1, -1}, errors.New("no such element")
}

func (b *AbstractBoard) FindAll(symbol Element) []Point {
	var points []Point
	for i, el := range b.Content {
		if el == rune(symbol) {
			points = append(points, b.IndexToPoint(i))
		}
	}
	return points
}

func (b *AbstractBoard) FindAllOf(symbols []Element) []Point {
	var points []Point
	for _, symbol := range symbols {
		for _, p := range b.FindAll(symbol) {
			points = append(points, p)
		}
	}
	return points
}

func (b *AbstractBoard) IsAt(point Point, element Element) bool {
	index := b.PointToIndex(point)
	if index < 0 || index >= len(b.Content) {
		return false
	}
	return b.Content[index] == rune(element)
}

func (b *AbstractBoard) IsAtAny(point Point, element []Element) bool {
	for _, el := range element {
		if b.IsAt(point, el) {
			return true
		}
	}
	return false
}

func (b *AbstractBoard) GetAt(p Point) (Element, error) {
	if p.X < 0 || p.X > b.Size {
		return Element(' '), errors.New(fmt.Sprintf("invalid x value: %d", p.X))
	}
	if p.Y < 0 || p.Y > b.Size {
		return Element(' '), errors.New(fmt.Sprintf("invalid y value: %d", p.Y))
	}
	return Element(b.Content[b.PointToIndex(p)]), nil
}

func (b *AbstractBoard) IsNear(p Point, element Element) bool {
	return b.IsAt(Point{p.X, p.Y - 1}, element) ||
		b.IsAt(Point{p.X, p.Y + 1}, element) ||
		b.IsAt(Point{p.X - 1, p.Y}, element) ||
		b.IsAt(Point{p.X + 1, p.Y}, element)
}

func (b *AbstractBoard) CountNear(p Point, element Element) int {
	counter := 0
	if b.IsAt(Point{p.X, p.Y}, element) {
		counter++
	}
	if b.IsAt(Point{p.X, p.Y - 1}, element) {
		counter++
	}
	if b.IsAt(Point{p.X, p.Y + 1}, element) {
		counter++
	}
	if b.IsAt(Point{p.X - 1, p.Y}, element) {
		counter++
	}
	if b.IsAt(Point{p.X + 1, p.Y}, element) {
		counter++
	}
	return counter
}
