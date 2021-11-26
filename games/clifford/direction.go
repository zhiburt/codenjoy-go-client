package clifford

import (
	"github.com/codenjoyme/codenjoy-go-client/engine/direction"
)

const (
	left                          = direction.Left
	right                         = direction.Right
	up                            = direction.Up
	down                          = direction.Down
	stop                          = direction.Stop
	crackLeft      direction.Base = "ACT,LEFT"
	crackRight     direction.Base = "ACT,RIGHT"
	die            direction.Base = "ACT(0)"
	shootLeft      direction.Base = "ACT(1),LEFT"
	shootRight     direction.Base = "ACT(1),RIGHT"
	openDoorLeft   direction.Base = "ACT(2),LEFT"
	openDoorRight  direction.Base = "ACT(2),RIGHT"
	closeDoorLeft  direction.Base = "ACT(3),LEFT"
	closeDoorRight direction.Base = "ACT(3),RIGHT"
)

var directions = direction.Map{
	left:           direction.New(0, -1, 0, direction.Left), // move
	right:          direction.New(1, 1, 0, direction.Right), // move
	up:             direction.New(2, 0, -1, direction.Up),   // move
	down:           direction.New(3, 0, 1, direction.Down),  // move
	stop:           direction.New(6, 0, 0, ""),              // stay
	crackLeft:      direction.New(4, 0, 0, crackLeft),       // crack ground at left
	crackRight:     direction.New(5, 0, 0, crackRight),      // crack ground at right
	die:            direction.New(7, 0, 0, die),             // suicide
	shootLeft:      direction.New(8, 0, 0, shootLeft),       // shoot to the left
	shootRight:     direction.New(9, 0, 0, shootRight),      // shoot to the right
	openDoorLeft:   direction.New(10, 0, 0, openDoorLeft),   // open door on left
	openDoorRight:  direction.New(11, 0, 0, openDoorRight),  // open door on right
	closeDoorLeft:  direction.New(12, 0, 0, closeDoorLeft),  // close door on left
	closeDoorRight: direction.New(13, 0, 0, closeDoorRight), // close door on right
}
