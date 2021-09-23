package clifford

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
		Elements["HERO_LEFT"],
		Elements["HERO_RIGHT"],
		Elements["HERO_CRACK_LEFT"],
		Elements["HERO_CRACK_RIGHT"],
		Elements["HERO_LADDER"],
		Elements["HERO_FALL_LEFT"],
		Elements["HERO_FALL_RIGHT"],
		Elements["HERO_PIPE_LEFT"],
		Elements["HERO_PIPE_RIGHT"],
		Elements["HERO_DIE"],
		Elements["HERO_PIT_LEFT"],
		Elements["HERO_PIT_RIGHT"],

		Elements["HERO_MASK_LEFT"],
		Elements["HERO_MASK_RIGHT"],
		Elements["HERO_MASK_CRACK_LEFT"],
		Elements["HERO_MASK_CRACK_RIGHT"],
		Elements["HERO_MASK_LADDER"],
		Elements["HERO_MASK_FALL_LEFT"],
		Elements["HERO_MASK_FALL_RIGHT"],
		Elements["HERO_MASK_PIPE_LEFT"],
		Elements["HERO_MASK_PIPE_RIGHT"],
		Elements["HERO_MASK_DIE"],
		Elements["HERO_MASK_PIT_LEFT"],
		Elements["HERO_MASK_PIT_RIGHT"])
	if len(points) == 0 {
		panic("hero element has not been found")
	}
	return points[0]
}

func (b *Board) FindOtherHeroes() []*engine.Point {
	return b.board.Find(
		Elements["OTHER_HERO_DIE"],
		Elements["OTHER_HERO_CRACK_LEFT"],
		Elements["OTHER_HERO_CRACK_RIGHT"],
		Elements["OTHER_HERO_LADDER"],
		Elements["OTHER_HERO_LEFT"],
		Elements["OTHER_HERO_RIGHT"],
		Elements["OTHER_HERO_FALL_LEFT"],
		Elements["OTHER_HERO_FALL_RIGHT"],
		Elements["OTHER_HERO_PIPE_LEFT"],
		Elements["OTHER_HERO_PIPE_RIGHT"],
		Elements["OTHER_HERO_PIT_LEFT"],
		Elements["OTHER_HERO_PIT_RIGHT"],

		Elements["OTHER_HERO_MASK_DIE"],
		Elements["OTHER_HERO_MASK_CRACK_LEFT"],
		Elements["OTHER_HERO_MASK_CRACK_RIGHT"],
		Elements["OTHER_HERO_MASK_LADDER"],
		Elements["OTHER_HERO_MASK_LEFT"],
		Elements["OTHER_HERO_MASK_RIGHT"],
		Elements["OTHER_HERO_MASK_FALL_LEFT"],
		Elements["OTHER_HERO_MASK_FALL_RIGHT"],
		Elements["OTHER_HERO_MASK_PIPE_LEFT"],
		Elements["OTHER_HERO_MASK_PIPE_RIGHT"],
		Elements["OTHER_HERO_MASK_PIT_LEFT"],
		Elements["OTHER_HERO_MASK_PIT_RIGHT"])
}

func (b *Board) FindEnemyHeroes() []*engine.Point {
	return b.board.Find(
		Elements["ENEMY_HERO_DIE"],
		Elements["ENEMY_HERO_CRACK_LEFT"],
		Elements["ENEMY_HERO_CRACK_RIGHT"],
		Elements["ENEMY_HERO_LADDER"],
		Elements["ENEMY_HERO_LEFT"],
		Elements["ENEMY_HERO_RIGHT"],
		Elements["ENEMY_HERO_FALL_LEFT"],
		Elements["ENEMY_HERO_FALL_RIGHT"],
		Elements["ENEMY_HERO_PIPE_LEFT"],
		Elements["ENEMY_HERO_PIPE_RIGHT"],
		Elements["ENEMY_HERO_PIT_LEFT"],
		Elements["ENEMY_HERO_PIT_RIGHT"],

		Elements["ENEMY_HERO_MASK_DIE"],
		Elements["ENEMY_HERO_MASK_CRACK_LEFT"],
		Elements["ENEMY_HERO_MASK_CRACK_RIGHT"],
		Elements["ENEMY_HERO_MASK_LADDER"],
		Elements["ENEMY_HERO_MASK_LEFT"],
		Elements["ENEMY_HERO_MASK_RIGHT"],
		Elements["ENEMY_HERO_MASK_FALL_LEFT"],
		Elements["ENEMY_HERO_MASK_FALL_RIGHT"],
		Elements["ENEMY_HERO_MASK_PIPE_LEFT"],
		Elements["ENEMY_HERO_MASK_PIPE_RIGHT"],
		Elements["ENEMY_HERO_MASK_PIT_LEFT"],
		Elements["ENEMY_HERO_MASK_PIT_RIGHT"])
}

func (b *Board) FindRobbers() []*engine.Point {
	return b.board.Find(
		Elements["ROBBER_LADDER"],
		Elements["ROBBER_LEFT"],
		Elements["ROBBER_RIGHT"],
		Elements["ROBBER_FALL_LEFT"],
		Elements["ROBBER_FALL_RIGHT"],
		Elements["ROBBER_PIPE_LEFT"],
		Elements["ROBBER_PIPE_RIGHT"],
		Elements["ROBBER_PIT_LEFT"],
		Elements["ROBBER_PIT_RIGHT"])
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
