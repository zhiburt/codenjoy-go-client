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

package verland

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

func directions() (direction.Map, error) {
	return direction.NewMap(
		direction.New(0, -1, 0, left),           // move
		direction.New(1, 1, 0, right),           // move
		direction.New(2, 0, -1, up),             // move
		direction.New(3, 0, 1, down),            // move
		direction.New(6, 0, 0, stop),            // stay
		direction.New(4, 0, 0, crackLeft),       // crack ground at left
		direction.New(5, 0, 0, crackRight),      // crack ground at right
		direction.New(7, 0, 0, die),             // suicide
		direction.New(8, 0, 0, shootLeft),       // shoot to the left
		direction.New(9, 0, 0, shootRight),      // shoot to the right
		direction.New(10, 0, 0, openDoorLeft),   // open door on left
		direction.New(11, 0, 0, openDoorRight),  // open door on right
		direction.New(12, 0, 0, closeDoorLeft),  // close door on left
		direction.New(13, 0, 0, closeDoorRight), // close door on right
	)
}
