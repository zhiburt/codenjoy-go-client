package mollymage

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
	"math"
	"math/rand"
	"time"

	"github.com/codenjoyme/codenjoy-go-client/direction"
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
	actTimer uint
}

func NewSolver() engine.Solver {
	return &Solver{}
}

func (solver *Solver) Answer(message string) string {
	board := newBoard(message)
	fmt.Println("Board \n" + board.String())
	action := nextAction(solver, board)
	fmt.Println("\nAnswer: " + action.String())
	fmt.Println("-------------------------------------------------------------")
	return action.String()
}

func nextAction(solver *Solver, b *board) direction.Direction {
	// Algorithm is based on graph A* algorithm.
	//
	// We don't care about other players,
	// We just try to not get on BOMBs and
	// place our bombs as often as possible and as close
	// as possible to chests
	//
	// we could actually estimate how much people in the area so
	// if there's a lot of chests but there's a lot of people there
	// we would chose a nearest chest as a target
	//
	// don't forget that chests by which side BOMBs are spawned has no value any more

	estimationFunction(b)

	if solver.actTimer == 0 {
		solver.actTimer = 3
		return directions.Get(act)
	}

	solver.actTimer -= 1
	for i := 0; i < 4; i += 1 {
		nextMove := randomMove()
		nextPos := doMove(b.findHero(), &nextMove)
		if !inDanger(b, nextPos) {
			return nextMove
		}
	}

	if !inDanger(b, b.findHero()) {
		return directions.Get(stop)
	}

	return directions.Get(act)
}

func estimationFunction(b *board) int32 {
	if inDanger(b, b.findHero()) {
		return math.MinInt32
	}

	return 0
}

func inDanger(b *board, hero *engine.Point) bool {
	blasts := b.findBlasts()
	for _, blast := range blasts {
		if blast == hero {
			return true
		}
	}

	return false
}

func isIn(pos *engine.Point, arr ...*engine.Point) bool {
	for _, point := range arr {
		if point == pos {
			return true
		}
	}

	return false
}

func randomMove() direction.Direction {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4)
	switch n {
	case 0:
		return directions.Get(up)
	case 1:
		return directions.Get(down)
	case 2:
		return directions.Get(left)
	default:
		return directions.Get(right)
	}
}

func doMove(pos *engine.Point, direction *direction.Direction) *engine.Point {
	return engine.NewPoint(direction.ChangeX(pos.X()), direction.ChangeY(pos.Y()))
}
