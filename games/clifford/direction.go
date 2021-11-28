package clifford

import (
	"github.com/codenjoyme/codenjoy-go-client/direction"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

const (
	left           engine.Action = "LEFT"
	right          engine.Action = "RIGHT"
	up             engine.Action = "UP"
	down           engine.Action = "DOWN"
	stop           engine.Action = "STOP"
	crackLeft      engine.Action = "ACT,LEFT"
	crackRight     engine.Action = "ACT,RIGHT"
	die            engine.Action = "ACT(0)"
	shootLeft      engine.Action = "ACT(1),LEFT"
	shootRight     engine.Action = "ACT(1),RIGHT"
	openDoorLeft   engine.Action = "ACT(2),LEFT"
	openDoorRight  engine.Action = "ACT(2),RIGHT"
	closeDoorLeft  engine.Action = "ACT(3),LEFT"
	closeDoorRight engine.Action = "ACT(3),RIGHT"
)

var directions = direction.Directions{
	direction.Left:  engine.NewPoint(-1, 0), // move
	direction.Right: engine.NewPoint(1, 0),  // move
	direction.Up:    engine.NewPoint(0, -1), // move
	direction.Down:  engine.NewPoint(0, 1),  // move
	direction.Stop:  engine.NewPoint(0, 0),  // stay
}
