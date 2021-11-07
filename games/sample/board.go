package sample

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
    "reflect"
    "sort"
)

type Board struct {
    board *engine.GameBoard
}

func NewBoard(message string) *Board {
    ElementsValues := make([]rune, 0, len(Elements))
    for _, e := range Elements {
        ElementsValues = append(ElementsValues, e)
    }
    return &Board{engine.NewGameBoard(ElementsValues, message)}
}

func (b *Board) GetAt(pt *engine.Point) rune {
    if !pt.IsValid(b.board.GetSize()) {
        return Elements["WALL"]
    }
    return b.board.GetAt(pt)
}

func (b *Board) FindHero() *engine.Point {
    points := b.board.Find(
        Elements["HERO"],
        Elements["DEAD_HERO"])

    if len(points) == 0 {
        panic("Hero element has not been found")
    }
    return points[0]
}

func (b *Board) IsGameOver() bool {
    return len(b.board.Find(Elements["DEAD_HERO"])) != 0
}

func (b *Board) FindOtherHeroes() []*engine.Point {
    return b.board.Find(
        Elements["OTHER_HERO"],
        Elements["OTHER_DEAD_HERO"])
}

func (b *Board) FindBarriers() []*engine.Point {
    var points []*engine.Point
    points = appendIfMissing(points, b.FindWalls()...)
    points = appendIfMissing(points, b.FindBombs()...)
    points = appendIfMissing(points, b.FindOtherHeroes()...)
    sort.Sort(engine.SortedPoints(points))
    return points
}

func appendIfMissing(slice []*engine.Point, points ...*engine.Point) []*engine.Point {
    for _, p := range points {
        existed := false
        for _, ele := range slice {
            if reflect.DeepEqual(ele, p) {
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

func (b *Board) FindWalls() []*engine.Point {
    return b.board.Find(Elements["WALL"])
}

func (b *Board) FindBombs() []*engine.Point {
    return b.board.Find(Elements["BOMB"])
}

func (b *Board) FindGold() []*engine.Point {
    return b.board.Find(Elements["GOLD"])
}

func (b *Board) String() string {
    return b.board.String() +
        "\nHero at: " + b.FindHero().String() +
        "\nOther heroes at: " + fmt.Sprintf("%v", b.FindOtherHeroes()) +
        "\nBombs at: " + fmt.Sprintf("%v", b.FindBombs()) +
        "\nGold at: " + fmt.Sprintf("%v", b.FindGold())
}