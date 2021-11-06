package clifford

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

func (b *Board) IsGameOver() bool {
    return len(b.board.Find(Elements["HERO_DIE"], Elements["HERO_MASK_DIE"])) != 0
}

func (b *Board) FindHero() *engine.Point {
    points := b.board.Find(
        Elements["HERO_DIE"],
        Elements["HERO_LADDER"],
        Elements["HERO_LEFT"],
        Elements["HERO_RIGHT"],
        Elements["HERO_FALL"],
        Elements["HERO_PIPE"],
        Elements["HERO_PIT"],

        Elements["HERO_MASK_DIE"],
        Elements["HERO_MASK_LADDER"],
        Elements["HERO_MASK_LEFT"],
        Elements["HERO_MASK_RIGHT"],
        Elements["HERO_MASK_FALL"],
        Elements["HERO_MASK_PIPE"],
        Elements["HERO_MASK_PIT"])

    if len(points) == 0 {
        panic("hero element has not been found")
    }
    return points[0]
}

func (b *Board) FindOtherHeroes() []*engine.Point {
    return b.board.Find(
        Elements["OTHER_HERO_DIE"],
        Elements["OTHER_HERO_LADDER"],
        Elements["OTHER_HERO_LEFT"],
        Elements["OTHER_HERO_RIGHT"],
        Elements["OTHER_HERO_FALL"],
        Elements["OTHER_HERO_PIPE"],
        Elements["OTHER_HERO_PIT"],

        Elements["OTHER_HERO_MASK_DIE"],
        Elements["OTHER_HERO_MASK_LADDER"],
        Elements["OTHER_HERO_MASK_LEFT"],
        Elements["OTHER_HERO_MASK_RIGHT"],
        Elements["OTHER_HERO_MASK_FALL"],
        Elements["OTHER_HERO_MASK_PIPE"],
        Elements["OTHER_HERO_MASK_PIT"])
}

func (b *Board) FindEnemyHeroes() []*engine.Point {
    return b.board.Find(
        Elements["ENEMY_HERO_DIE"],
        Elements["ENEMY_HERO_LADDER"],
        Elements["ENEMY_HERO_LEFT"],
        Elements["ENEMY_HERO_RIGHT"],
        Elements["ENEMY_HERO_FALL"],
        Elements["ENEMY_HERO_PIPE"],
        Elements["ENEMY_HERO_PIT"],

        Elements["ENEMY_HERO_MASK_DIE"],
        Elements["ENEMY_HERO_MASK_LADDER"],
        Elements["ENEMY_HERO_MASK_LEFT"],
        Elements["ENEMY_HERO_MASK_RIGHT"],
        Elements["ENEMY_HERO_MASK_FALL"],
        Elements["ENEMY_HERO_MASK_PIPE"],
        Elements["ENEMY_HERO_MASK_PIT"])
}

func (b *Board) FindRobbers() []*engine.Point {
    return b.board.Find(
        Elements["ROBBER_LADDER"],
        Elements["ROBBER_LEFT"],
        Elements["ROBBER_RIGHT"],
        Elements["ROBBER_FALL"],
        Elements["ROBBER_PIPE"],
        Elements["ROBBER_PIT"])
}

func (b *Board) FindBarriers() []*engine.Point {
    return b.board.Find(
        Elements["BRICK"],
        Elements["STONE"])
}

func (b *Board) FindPits() []*engine.Point {
    return b.board.Find(
        Elements["CRACK_PIT"],
        Elements["PIT_FILL_1"],
        Elements["PIT_FILL_2"],
        Elements["PIT_FILL_3"],
        Elements["PIT_FILL_4"])
}

func (b *Board) FindClues() []*engine.Point {
    return b.board.Find(
        Elements["CLUE_KNIFE"],
        Elements["CLUE_GLOVE"],
        Elements["CLUE_RING"])
}

func (b *Board) FindBackways() []*engine.Point {
    return b.board.Find(Elements["BACKWAY"])
}

func (b *Board) FindPotions() []*engine.Point {
    return b.board.Find(Elements["MASK_POTION"])
}

func (b *Board) FindDoors() []*engine.Point {
    return b.board.Find(        
        Elements["OPENED_DOOR_GOLD"],
        Elements["OPENED_DOOR_SILVER"],
        Elements["OPENED_DOOR_BRONZE"],
        Elements["CLOSED_DOOR_GOLD"],
        Elements["CLOSED_DOOR_SILVER"],
        Elements["CLOSED_DOOR_BRONZE"])
}

func (b *Board) FindKeys() []*engine.Point {
    return b.board.Find(
        Elements["KEY_GOLD"],
        Elements["KEY_SILVER"],
        Elements["KEY_BRONZE"])
}

func (b *Board) String() string {
    return b.board.String() +
        "\nHero at: " + b.FindHero().String() +
        "\nOther heroes at: " + fmt.Sprintf("%v", b.FindOtherHeroes()) +
        "\nEnemy heroes at: " + fmt.Sprintf("%v", b.FindEnemyHeroes()) +
        "\nRobbers at: " + fmt.Sprintf("%v", b.FindRobbers()) +
        "\nMask potions at: " + fmt.Sprintf("%v", b.FindPotions()) +
        "\nKeys at: " + fmt.Sprintf("%v", b.FindKeys())
}
