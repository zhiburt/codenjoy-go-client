package mollymage

import (
	"github.com/codenjoyme/codenjoy-go-client/engine/direction"
)

const (
	left                 = direction.Left
	right                = direction.Right
	up                   = direction.Up
	down                 = direction.Down
	stop                 = direction.Stop
	act   direction.Base = "ACT"
)

func directions() (direction.Map, error) {
	return direction.NewMap(
		direction.New(1, -1, 0, left), // move
		direction.New(2, 1, 0, right), // move
		direction.New(3, 0, 1, up),    // move
		direction.New(4, 0, -1, down), // move
		direction.New(0, 0, 0, stop),  // stay
		direction.New(5, 0, 0, act),   // act
	)
}
