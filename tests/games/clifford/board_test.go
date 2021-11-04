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
    "github.com/codenjoyme/codenjoy-go-client/games/clifford"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsGameOver(t *testing.T) {
    board := clifford.NewBoard("###" + "##►" + "###")
    assert.Equal(t, false, board.IsGameOver())

    board = clifford.NewBoard("###" + "Ѡ##" + "###")
    assert.Equal(t, true, board.IsGameOver())
    board = clifford.NewBoard("###" + "x##" + "###")
    assert.Equal(t, true, board.IsGameOver())
}

func TestFindHeroNoResult(t *testing.T) {
    board := clifford.NewBoard("###" + "###" + "###")
    assert.Panics(t, func() { fmt.Sprintf("%v", board.FindHero()) })
}

func TestFindHero(t *testing.T) {
    board := clifford.NewBoard("Ѡ##" + "###" + "###")
    assert.Equal(t, "[0,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("#Я#" + "###" + "###")
    assert.Equal(t, "[1,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("##R" + "###" + "###")
    assert.Equal(t, "[2,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "Y##" + "###")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "#◄#" + "###")
    assert.Equal(t, "[1,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "##►" + "###")
    assert.Equal(t, "[2,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "]##")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "#[#")
    assert.Equal(t, "[1,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##{")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##}")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##⍃")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##⍄")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))

    board = clifford.NewBoard("⍃⍄ѠЯ" + "RY◄►" + "][{}" + "####")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindHero()))
}

func TestFindHero_Mask(t *testing.T) {
    board := clifford.NewBoard("x##" + "###" + "###")
    assert.Equal(t, "[0,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("#⊰#" + "###" + "###")
    assert.Equal(t, "[1,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("##⊱" + "###" + "###")
    assert.Equal(t, "[2,2]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "⍬##" + "###")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "#⊲#" + "###")
    assert.Equal(t, "[1,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "##⊳" + "###")
    assert.Equal(t, "[2,1]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "⊅##")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "#⊄#")
    assert.Equal(t, "[1,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##⋜")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##⋝")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##ᐊ")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))
    board = clifford.NewBoard("###" + "###" + "##ᐅ")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.FindHero()))

    board = clifford.NewBoard("ᐊᐅx⊰" + "⊱⍬⊲⊳" + "⊅⊄⋜⋝" + "####")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindHero()))
}

func TestFindOtherHeroes(t *testing.T) {
    board := clifford.NewBoard("##Z⌋" + "⌊U)(" + "⊐⊏ЭЄ" + "ᗉᗆ##")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,1] "+
        "[2,2] [2,3] [3,1] [3,2] [3,3]]", fmt.Sprintf("%v", board.FindOtherHeroes()))

    board = clifford.NewBoard("##⋈⋰" + "⋱⋕⋊⋉" + "⋣⋢⊣⊢" + "ᗏᗌ##")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,1] "+
        "[2,2] [2,3] [3,1] [3,2] [3,3]]", fmt.Sprintf("%v", board.FindOtherHeroes()))
}
func TestFindEnemyHeroes(t *testing.T) {
    board := clifford.NewBoard("##Ž⟧" + "⟦Ǔ❫❪" + "⋥⋤ǮĚ" + "⇇⇉##")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,1] "+
        "[2,2] [2,3] [3,1] [3,2] [3,3]]", fmt.Sprintf("%v", board.FindEnemyHeroes()))

    board = clifford.NewBoard("##⧓⇢" + "⇠≠⧒⧑" + "⌫⌦❵❴" + "⬱⇶##")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,1] "+
        "[2,2] [2,3] [3,1] [3,2] [3,3]]", fmt.Sprintf("%v", board.FindEnemyHeroes()))
}

func TestFindRobbers(t *testing.T) {
    board := clifford.NewBoard("Q«»" + "‹›<" + ">⍇⍈")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] "+
        "[1,1] [1,2] [2,0] [2,1] [2,2]]", fmt.Sprintf("%v", board.FindRobbers()))
}

func TestFindBarriers(t *testing.T) {
    board := clifford.NewBoard("  #" + "  ☼" + "   ")
    assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.FindBarriers()))
}

func TestFindPits(t *testing.T) {
    board := clifford.NewBoard("123" + "4**" + "###")
    assert.Equal(t, "[[0,1] [0,2] [1,1] [1,2] [2,1] [2,2]]", fmt.Sprintf("%v", board.FindPits()))
}

func TestFindClues(t *testing.T) {
    board := clifford.NewBoard("##$" + "##&" + "##@")
    assert.Equal(t, "[[2,0] [2,1] [2,2]]", fmt.Sprintf("%v", board.FindClues()))
}

func TestFindBackways(t *testing.T) {
    board := clifford.NewBoard("##⊛" + "###" + "###")
    assert.Equal(t, "[[2,2]]", fmt.Sprintf("%v", board.FindBackways()))
}

func TestFindPotions(t *testing.T) {
    board := clifford.NewBoard("##S" + "###" + "###")
    assert.Equal(t, "[[2,2]]", fmt.Sprintf("%v", board.FindPotions()))
}

func TestFindDoors(t *testing.T) {
    board := clifford.NewBoard("⍙⍚⍜" + "⍍⌺⌼" + "###")
    assert.Equal(t, "[[0,1] [0,2] [1,1] [1,2] [2,1] [2,2]]", fmt.Sprintf("%v", board.FindDoors()))
}

func TestFindKeys(t *testing.T) {
    board := clifford.NewBoard("✦✼⍟" + "###" + "###")
    assert.Equal(t, "[[0,2] [1,2] [2,2]]", fmt.Sprintf("%v", board.FindKeys()))
}

func TestReport(t *testing.T) {
    board := clifford.NewBoard("board=" +
        "☼☼☼☼☼☼☼☼☼" +
        "☼ ►*## $☼" +
        "☼ H ⧒⧒ ✼☼" +
        "☼ H  1 ⍍☼" +
        "☼S ⊐ &  ☼" +
        "☼ ✦ ~~~ ☼" +
        "☼Z3 ⌺ ⊏ ☼" +
        "☼ @@ ⍈Q ☼" +
        "☼☼☼☼☼☼☼☼☼")
    assert.Equal(t, ""+
        "☼☼☼☼☼☼☼☼☼\n" +
        "☼ ►*## $☼\n" +
        "☼ H ⧒⧒ ✼☼\n" +
        "☼ H  1 ⍍☼\n" +
        "☼S ⊐ &  ☼\n" +
        "☼ ✦ ~~~ ☼\n" +
        "☼Z3 ⌺ ⊏ ☼\n" +
        "☼ @@ ⍈Q ☼\n" +
        "☼☼☼☼☼☼☼☼☼\n"+
        "\n"+
        "Hero at: [2,7]\n"+
        "Other heroes at: [[1,2] [3,4] [6,2]]\n"+
        "Enemy heroes at: [[4,6] [5,6]]\n"+
        "Robbers at: [[5,1] [6,1]]\n"+
        "Mask potions at: [[1,4]]\n" +
        "Keys at: [[2,3] [7,6]]", fmt.Sprintf("%v", board.String()))
}
