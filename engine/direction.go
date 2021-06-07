package engine

type Direction int

const (
    UP Direction = iota
    DOWN
    RIGHT
    LEFT
)

func (d Direction) String() string {
    return [...]string{"UP", "DOWN", "RIGHT", "LEFT"}[d]
}
