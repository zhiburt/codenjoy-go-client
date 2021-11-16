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
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
}

func NewSolver() engine.Solver {
	return Solver{}
}

func (Solver) Answer(message string) string {
	board := newBoard(message)
	fmt.Println("Board \n" + board.String())
	action := nextAction(board)
	fmt.Println("\nAnswer: " + action.String())
	fmt.Println("-------------------------------------------------------------")
	return action.String()
}

func nextAction(b *board) engine.Direction {
	// TODO: write your code here
	return engine.ACT
}
