package mollymage

import (
    "fmt"
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/codenjoyme/codenjoy-go-client/games/mollymage"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetAtInvalidPoint(t *testing.T) {
    board := mollymage.NewBoard("###" + "###" + "###")
    assert.Equal(t, mollymage.Elements["WALL"], board.GetAt(engine.NewPoint(-1, -1)))
}

func TestFindHero(t *testing.T) {
    board := mollymage.NewBoard("#☺#" + "###" + "###")
    assert.Equal(t, "[1,2]", board.FindHero().String())

    board = mollymage.NewBoard("###" + "#☻#" + "###")
    assert.Equal(t, "[1,1]", board.FindHero().String())

    board = mollymage.NewBoard("###" + "###" + "#Ѡ#")
    assert.Equal(t, "[1,0]", board.FindHero().String())

    board = mollymage.NewBoard("Ѡ☺☻" + "###" + "###")
    assert.Equal(t, "[0,2]", board.FindHero().String())
}

func TestFindHeroNoResult(t *testing.T) {
    board := mollymage.NewBoard("###" + "###" + "###")
    assert.Panics(t, func() { board.FindHero() })
}

func TestIsGameOver(t *testing.T) {
    board := mollymage.NewBoard("###" + "##☺" + "###")
    assert.Equal(t, false, board.IsGameOver())

    board = mollymage.NewBoard("###" + "Ѡ##" + "###")
    assert.Equal(t, true, board.IsGameOver())
}

func TestFindOtherHeroes(t *testing.T) {
    board := mollymage.NewBoard("#♥#" + "#♠#" + "#♣#")
    assert.Equal(t, "[[1,0] [1,1] [1,2]]", fmt.Sprintf("%v", board.FindOtherHeroes()))
}

func TestFindEnemyHeroes(t *testing.T) {
    board := mollymage.NewBoard("#♡#" + "#♤#" + "#♧#")
    assert.Equal(t, "[[1,0] [1,1] [1,2]]", fmt.Sprintf("%v", board.FindEnemyHeroes()))
}

func TestFindBarriers(t *testing.T) {
    board := mollymage.NewBoard("☼&#" + "123" + "♥♠♣")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0] [2,1] [2,2]]",
        fmt.Sprintf("%v", board.FindBarriers()))
}

func TestFindWalls(t *testing.T) {
    board := mollymage.NewBoard("###" + "☼##" + "☼##")
    assert.Equal(t, "[[0,0] [0,1]]", fmt.Sprintf("%v", board.FindWalls()))
}

func TestFindGhosts(t *testing.T) {
    board := mollymage.NewBoard("##&" + "##&" + "###")
    assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.FindGhosts()))
}

func TestFindTreasureBoxes(t *testing.T) {
    board := mollymage.NewBoard("҉#҉" + "҉҉҉" + "҉#҉")
    assert.Equal(t, "[[1,0] [1,2]]", fmt.Sprintf("%v", board.FindTreasureBoxes()))
}

func TestFindPotions(t *testing.T) {
    board := mollymage.NewBoard("123" + "45#" + "☻♠#")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,2]]",
        fmt.Sprintf("%v", board.FindPotions()))
}

func TestFindBlasts(t *testing.T) {
    board := mollymage.NewBoard("###" + "###" + "##҉")
    assert.Equal(t, "[[2,0]]", fmt.Sprintf("%v", board.FindBlasts()))
}

func TestFindPerks(t *testing.T) {
    board := mollymage.NewBoard("#cr" + "#i+" + "#TA")
    assert.Equal(t, "[[1,0] [1,1] [1,2] [2,0] [2,1] [2,2]]", fmt.Sprintf("%v", board.FindPerks()))
}

func TestReport(t *testing.T) {
    board := mollymage.NewBoard("board=" +
        "☼☼☼☼☼☼☼☼☼" +
        "☼1 ♣   ♠☼" +
        "☼#2  &  ☼" +
        "☼# 3 ♣ ♠☼" +
        "☼☺  4   ☼" +
        "☼   ♡ H☻☼" +
        "☼x H ҉҉҉☼" +
        "☼& &    ☼" +
        "☼☼☼☼☼☼☼☼☼")
    assert.Equal(t, "☼☼☼☼☼☼☼☼☼\n"+
        "☼1 ♣   ♠☼\n"+
        "☼#2  &  ☼\n"+
        "☼# 3 ♣ ♠☼\n"+
        "☼☺  4   ☼\n"+
        "☼   ♡ H☻☼\n"+
        "☼x H ҉҉҉☼\n"+
        "☼& &    ☼\n"+
        "☼☼☼☼☼☼☼☼☼\n"+
        "\n"+
        "Hero at: [1,4]\n"+
        "Other heroes at: [[3,7] [5,5] [7,5] [7,7]]\n"+
        "Enemy heroes at: [[4,3]]\n"+
        "Ghosts at: [[1,1] [3,1] [5,6]]\n"+
        "Potions at: [[1,7] [2,6] [3,5] [4,4] [7,3] [7,5] [7,7]]\n"+
        "Blasts at: [[5,2] [6,2] [7,2]]\n"+
        "Expected blasts at: [[2,7]]", board.String())
}
