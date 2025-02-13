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
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAtInvalidPoint(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, elements["WALL"], board.getAt(engine.NewPoint(-1, -1)))
}

func TestFindHero(t *testing.T) {
	board := newBoard("☼☺☼" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, "[1,2]", board.findHero().String())

	board = newBoard("☼☼☼" + "☼☺☼" + "☼☼☼")
	assert.Equal(t, "[1,1]", board.findHero().String())

	board = newBoard("☼☼☼" + "☼☼☼" + "☼X☼")
	assert.Equal(t, "[1,0]", board.findHero().String())

	board = newBoard("X☺☻" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, "[0,2]", board.findHero().String())
}

func TestFindHeroNoResult(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼☼" + "☼☼☼")
	assert.Panics(t, func() { board.findHero() })
}

func TestIsGameOver(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼☺" + "☼☼☼")
	assert.Equal(t, false, board.isGameOver())

	board = newBoard("☼☼☼" + "X☼☼" + "☼☼☼")
	assert.Equal(t, true, board.isGameOver())
}

func TestFindOtherHeroes(t *testing.T) {
	board := newBoard("☼☻☼" + "☼Y☼" + "☼☻☼")
	assert.Equal(t, "[[1,0] [1,1] [1,2]]", fmt.Sprintf("%v", board.findOtherHeroes()))
}

func TestFindBarriers(t *testing.T) {
	board := newBoard("☼☼☼" + "xxx" + "☻☻☻")
	assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0] [2,1] [2,2]]",
		fmt.Sprintf("%v", board.findBarriers()))
}

func TestFindWalls(t *testing.T) {
	board := newBoard("   " + "☼  " + "☼  ")
	assert.Equal(t, "[[0,0] [0,1]]", fmt.Sprintf("%v", board.findWalls()))
}

func TestFindGold(t *testing.T) {
	board := newBoard("☼☼$" + "☼☼$" + "☼☼☼")
	assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.findGold()))
}

func TestFindBombs(t *testing.T) {
	board := newBoard("☼☼x" + "☼☼x" + "☼☼☼")
	assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.findBombs()))
}

func TestReport(t *testing.T) {
	board := newBoard("board=" +
		"☼☼☼☼☼☼☼☼☼" +
		"☼ x☺  Y ☼" +
		"☼  x    ☼" +
		"☼ $  ☻  ☼" +
		"☼      x☼" +
		"☼ ☻     ☼" +
		"☼       ☼" +
		"☼ $ ☻ x ☼" +
		"☼☼☼☼☼☼☼☼☼")
	assert.Equal(t,
		/*8*/ "☼☼☼☼☼☼☼☼☼\n"+
			/*7*/ "☼ x☺  Y ☼\n"+
			/*6*/ "☼  x    ☼\n"+
			/*5*/ "☼ $  ☻  ☼\n"+
			/*4*/ "☼      x☼\n"+
			/*3*/ "☼ ☻     ☼\n"+
			/*2*/ "☼       ☼\n"+
			/*1*/ "☼ $ ☻ x ☼\n"+
			/*0*/ "☼☼☼☼☼☼☼☼☼\n"+
			/*012345678*/
			"\n"+
			"Hero at: [3,7]\n"+
			"Other heroes at: [[2,3] [4,1] [5,5] [6,7]]\n"+
			"Bombs at: [[2,7] [3,6] [6,1] [7,4]]\n"+
			"Gold at: [[2,1] [2,5]]", board.String())
}
