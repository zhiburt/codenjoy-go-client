package engine

import "fmt"

type Direction struct {
	name  string
	value int
	dx    int
	dy    int
}

func (d *Direction) Value() int {
	return d.value
}

func (d *Direction) ChangeX(x int) int {
	return x + d.dx
}

func (d *Direction) ChangeY(y int) int {
	return y + d.dy
}

func (d *Direction) Inverted() Direction {
	if d.name == "LEFT" {
		return RIGHT
	}
	if d.name == "RIGHT" {
		return LEFT
	}
	if d.name == "UP" {
		return DOWN
	}
	if d.name == "DOWN" {
		return UP
	}
	panic(fmt.Sprintf("Cant invert for: %v", d))
}

func (d *Direction) String() string {
	return d.name
}

var LEFT = Direction{"LEFT", 0, -1, 0}
var RIGHT = Direction{"RIGHT", 1, 1, 0}
var UP = Direction{"UP", 2, 0, 1}
var DOWN = Direction{"DOWN", 3, 0, -1}
var ACT = Direction{"ACT", 4, 0, 0}
var STOP = Direction{"STOP", 5, 0, 0}
