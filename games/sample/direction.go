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

package sample

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

var directions = direction.Map{
	left:  direction.New(1, -1, 0, direction.Left), // move
	right: direction.New(2, 1, 0, direction.Right), // move
	up:    direction.New(3, 0, 1, direction.Up),    // move
	down:  direction.New(4, 0, -1, direction.Down), // move
	stop:  direction.New(0, 0, 0, ""),              // stay
	act:   direction.New(5, 0, 0, act),             // act
}
