package clifford

import (
	"github.com/codenjoyme/codenjoy-go-client/engine/direction"
)

type directions map[direction.Base]direction.Direction

const (
	crackLeft      = direction.Base("Act,Left")
	crackRight     = direction.Base("Act,Right")
	die            = direction.Base("Act(0)")
	shootLeft      = direction.Base("Act(1),Left")
	shootRight     = direction.Base("Act(1),Right")
	openDoorLeft   = direction.Base("Act(2),Left")
	openDoorRight  = direction.Base("Act(2),Right")
	closeDoorLeft  = direction.Base("Act(3),Left")
	closeDoorRight = direction.Base("Act(3),Right")
)

func initDirections() directions {
	m := make(directions)
	m[direction.Left] = direction.New(0, -1, 0, direction.Left)  // move
	m[direction.Right] = direction.New(1, 1, 0, direction.Right) // move
	m[direction.Up] = direction.New(2, 0, -1, direction.Up)      // move
	m[direction.Down] = direction.New(3, 0, 1, direction.Down)   // move
	m[crackLeft] = direction.New(4, 0, 0, crackLeft)             // crack ground at left
	m[crackRight] = direction.New(5, 0, 0, crackRight)           // crack ground at right
	m[direction.Stop] = direction.New(6, 0, 0, direction.Stop)   // stay
	m[die] = direction.New(7, 0, 0, die)                         // suicide
	m[shootLeft] = direction.New(8, 0, 0, shootLeft)             // shoot to the left
	m[shootRight] = direction.New(9, 0, 0, shootRight)           // shoot to the right
	m[openDoorLeft] = direction.New(10, 0, 0, openDoorLeft)      // open door on left
	m[openDoorRight] = direction.New(11, 0, 0, openDoorRight)    // open door on right
	m[closeDoorLeft] = direction.New(12, 0, 0, closeDoorLeft)    // close door on left
	m[closeDoorRight] = direction.New(13, 0, 0, closeDoorRight)  // close door on right

	return m
}
