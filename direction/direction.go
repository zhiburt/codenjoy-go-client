package direction

import (
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

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

type Directions map[Side]*engine.Point

type Side int

const (
	Left = iota
	Right
	Up
	Down
	Stop
)

func (side Side) Move(pos *engine.Point) *engine.Point {
	switch side {
	case Left:
		return engine.NewPoint(pos.X()-1, pos.Y())
	case Right:
		return engine.NewPoint(pos.X()+1, pos.Y())
	case Up:
		return engine.NewPoint(pos.X(), pos.Y()-1)
	case Down:
		return engine.NewPoint(pos.X()-1, pos.Y()+1)
	case Stop:
		return engine.NewPoint(pos.X(), pos.Y())
	default:
		panic("This point is unreachable")
	}
}

func FromPoint(pos *engine.Point) Directions {
	return Directions{
		Left:  engine.NewPoint(pos.X()-1, pos.Y()),
		Right: engine.NewPoint(pos.X()+1, pos.Y()),
		Up:    engine.NewPoint(pos.X(), pos.Y()-1),
		Down:  engine.NewPoint(pos.X(), pos.Y()+1),
		Stop:  engine.NewPoint(pos.X(), pos.Y()),
	}
}

func (m *Directions) Get(side Side) *engine.Point {
	return (*m)[side]
}

func (m *Directions) Inverted(side Side) Side {
	switch side {
	case Left:
		return Right
	case Right:
		return Left
	case Down:
		return Up
	case Up:
		return Down
	default:
		panic("This point is unreachable")
	}
}

func (m *Directions) Clockwise(side Side) Side {
	switch side {
	case Up:
		return Left
	case Left:
		return Down
	case Down:
		return Right
	case Right:
		return Up
	default:
		panic("This point is unreachable")
	}
}

func (m *Directions) ContrClockwise(side Side) Side {
	switch side {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("This point is unreachable")
	}
}

func (m *Directions) MirrorTopBottom(side Side) Side {
	switch side {
	case Up:
		return Left
	case Right:
		return Down
	case Down:
		return Right
	case Left:
		return Up
	default:
		panic("This point is unreachable")
	}
}

func (m *Directions) MirrorBottomTop(side Side) Side {
	switch side {
	case Up:
		return Right
	case Right:
		return Up
	case Down:
		return Left
	case Left:
		return Down
	default:
		panic("This point is unreachable")
	}
}
