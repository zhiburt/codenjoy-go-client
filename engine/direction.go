package engine

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
