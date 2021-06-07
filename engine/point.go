package engine

import "fmt"

type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

func (b *AbstractBoard) IsValid(p Point) bool {
    return (p.X >= 0 && p.X < b.Size) && (p.Y >= 0 && p.Y < b.Size)
}

func (b *AbstractBoard) IndexToPoint(index int) Point {
    return Point{
        X: index % b.Size,
        Y: (b.Size - (index / b.Size)) - 1,
    }
}

func (b *AbstractBoard) PointToIndex(p Point) int {
    index := (b.Size - 1 - p.Y) * b.Size
    index += p.X
    return index
}
