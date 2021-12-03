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
	"sort"

	"github.com/codenjoyme/codenjoy-go-client/engine"
)

const BLAST_RANGE int = 3

type board struct {
	board *engine.GameBoard
}

func newBoard(message string) *board {
	values := make([]rune, 0, len(elements))
	for _, e := range elements {
		values = append(values, e)
	}
	return &board{engine.NewGameBoard(values, message)}
}

func (b *board) getAt(pt *engine.Point) rune {
	if !pt.IsValid(b.board.Size()) {
		return elements["WALL"]
	}
	return b.board.GetAt(pt)
}

func (b *board) findHero() *engine.Point {
	points := b.board.Find(
		elements["HERO"],
		elements["POTION_HERO"],
		elements["DEAD_HERO"])
	if len(points) == 0 {
		panic("Hero element has not been found")
	}
	return points[0]
}

func (b *board) isGameOver() bool {
	return len(b.board.Find(elements["DEAD_HERO"])) != 0
}

func (b *board) findOtherHeroes() []*engine.Point {
	return b.board.Find(
		elements["OTHER_HERO"],
		elements["OTHER_POTION_HERO"],
		elements["OTHER_DEAD_HERO"])
}

func (b *board) findEnemyHeroes() []*engine.Point {
	return b.board.Find(
		elements["ENEMY_HERO"],
		elements["ENEMY_POTION_HERO"],
		elements["ENEMY_DEAD_HERO"])
}

func (b *board) findBarriers() []*engine.Point {
	var points []*engine.Point
	points = appendIfMissing(points, b.findWalls()...)
	points = appendIfMissing(points, b.findGhosts()...)
	points = appendIfMissing(points, b.findTreasureBoxes()...)
	points = appendIfMissing(points, b.findPotions()...)
	points = appendIfMissing(points, b.findOtherHeroes()...)
	points = appendIfMissing(points, b.findEnemyHeroes()...)
	sort.Sort(engine.SortedPoints(points))
	return points
}

func appendIfMissing(slice []*engine.Point, points ...*engine.Point) []*engine.Point {
	for _, p := range points {
		existed := false
		for _, ele := range slice {
			if ele.Equal(p) {
				existed = true
				break
			}
		}
		if !existed {
			slice = append(slice, p)
		}
	}
	return slice
}

func (b *board) findWalls() []*engine.Point {
	return b.board.Find(elements["WALL"])
}

func (b *board) findGhosts() []*engine.Point {
	return b.board.Find(elements["GHOST"])
}

func (b *board) findTreasureBoxes() []*engine.Point {
	return b.board.Find(elements["TREASURE_BOX"])
}

func (b *board) findPotions() []*engine.Point {
	return b.board.Find(
		elements["POTION_TIMER_1"],
		elements["POTION_TIMER_2"],
		elements["POTION_TIMER_3"],
		elements["POTION_TIMER_4"],
		elements["POTION_TIMER_5"],
		elements["POTION_HERO"],
		elements["OTHER_POTION_HERO"],
		elements["ENEMY_POTION_HERO"])
}

func (b *board) findBlasts() []*engine.Point {
	return b.board.Find(elements["BOOM"])
}

func (b *board) predictFutureBlasts() []*engine.Point {
	var barriers []*engine.Point
	for _, potion := range b.board.Find(elements["POTION_TIMER_1"]) {
		barriers = append(barriers, b.predictBlastsForOneSide(potion, engine.StepLeft)...)
		barriers = append(barriers, b.predictBlastsForOneSide(potion, engine.StepRight)...)
		barriers = append(barriers, b.predictBlastsForOneSide(potion, engine.StepUp)...)
		barriers = append(barriers, b.predictBlastsForOneSide(potion, engine.StepDown)...)
	}
	return barriers
}

type Move func(*engine.Point) *engine.Point

func (b *board) predictBlastsForOneSide(pt *engine.Point, nextStep Move) []*engine.Point {
	barriers := b.findBarriers()

	var points []*engine.Point
	for i := 0; i < BLAST_RANGE; i++ {
		pt = nextStep(pt)
		if !pt.IsValid(b.board.Size()) {
			break
		}
		isBarrier := false
		for _, barrier := range barriers {
			if barrier.Equal(pt) {
				isBarrier = true
				break
			}
		}
		if isBarrier == true {
			break
		}
		points = append(points, pt)
	}
	return points
}

func (b *board) findPerks() []*engine.Point {
	return b.board.Find(
		elements["POTION_COUNT_INCREASE"],
		elements["POTION_REMOTE_CONTROL"],
		elements["POTION_IMMUNE"],
		elements["POTION_BLAST_RADIUS_INCREASE"],
		elements["POISON_THROWER"],
		elements["POTION_EXPLODER"])
}

func (b *board) String() string {
	return b.board.String() +
		"\nHero at: " + b.findHero().String() +
		"\nOther heroes at: " + fmt.Sprintf("%v", b.findOtherHeroes()) +
		"\nEnemy heroes at: " + fmt.Sprintf("%v", b.findEnemyHeroes()) +
		"\nGhosts at: " + fmt.Sprintf("%v", b.findGhosts()) +
		"\nPotions at: " + fmt.Sprintf("%v", b.findPotions()) +
		"\nBlasts at: " + fmt.Sprintf("%v", b.findBlasts()) +
		"\nExpected blasts at: " + fmt.Sprintf("%v", b.predictFutureBlasts())
}

// todo:
// * blasts doesn't work it sometimes goes into its own blasts
// * devalue treasures which are about to be BOOMed
// * threasure look not for closest position but for a position which will affect most of the threasures
// * unction to determine when to plant the bomb the best
// * once we are by the threasure we are not afraid of potions && ghosts
// * add value to good potions

// * We can't go over potion so we need to fix it
// * Ghost on bomb can't be seen as bomb so adding a tracking of this cases could be cool
//   We could keep a map of blasts on the map and clean it when blasts appear
