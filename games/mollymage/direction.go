package mollymage

import (
	"github.com/codenjoyme/codenjoy-go-client/direction"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

const (
	left  engine.Action = "LEFT"
	right engine.Action = "RIGHT"
	up    engine.Action = "UP"
	down  engine.Action = "DOWN"
	stop  engine.Action = "STOP"
	act   engine.Action = "ACT"
)

var directions = direction.Directions{
	direction.Left:  engine.NewPoint(-1, 0), // move
	direction.Right: engine.NewPoint(1, 0),  // move
	direction.Up:    engine.NewPoint(0, -1), // move
	direction.Down:  engine.NewPoint(0, 1),  // move
	direction.Stop:  engine.NewPoint(0, 0),  // stay
}
