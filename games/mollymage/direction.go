package mollymage

import (
	"github.com/codenjoyme/codenjoy-go-client/direction"
)

const (
	left                 = direction.Left
	right                = direction.Right
	up                   = direction.Up
	down                 = direction.Down
	stop                 = direction.Stop
	act   direction.Base = "ACT"
)

var directions = direction.Map{
	left:  direction.New(-1, 0, direction.Left), // move
	right: direction.New(1, 0, direction.Right), // move
	up:    direction.New(0, 1, direction.Up),    // move
	down:  direction.New(0, -1, direction.Down), // move
	stop:  direction.New(0, 0, ""),              // stay
	act:   direction.New(0, 0, act),             // act
}
