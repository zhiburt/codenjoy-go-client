package direction

/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2021 Codenjoy
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this program.  If not, see
 * <http://www.gnu.org/licenses/gpl-3.0.html>.
 * #L%
 */

import (
	"fmt"
)

type Direction struct {
	name  Base
	value int
	dx    int
	dy    int
}

type Base string

const (
	Left  Base = "Left"
	Right Base = "Right"
	Up    Base = "Up"
	Down  Base = "Down"
	Act   Base = "Act"
	Stop  Base = "Stop"
)

func New(value, dx, dy int, name Base) Direction {
	return Direction{
		value: value,
		dx:    dx,
		dy:    dy,
		name:  name,
	}
}

func (d Direction) Value() int {
	return d.value
}

func (d Direction) ChangeX(x int) int {
	return x + d.dx
}

func (d Direction) ChangeY(y int) int {
	return y + d.dy
}

func (d Direction) Inverted() Base {
	switch d.name {
	case Left:
		return Right
	case Right:
		return Left
	case Down:
		return Up
	case Up:
		return Down
	default:
		panic(fmt.Sprintf("Cant invert for: %+v", d))
	}
}

func (d Direction) String() string {
	return string(d.name)
}

//var Left = Direction{"Left", 0, -1, 0}
//var Right = Direction{"Right", 1, 1, 0}
//var Up = Direction{"Up", 2, 0, 1}
//var Down = Direction{"Down", 3, 0, -1}
//var Act = Direction{"Act", 4, 0, 0}
//var Stop = Direction{"Stop", 5, 0, 0}
