package bomberman

import (
	"github.com/codenjoyme/codenjoy-go-client/engine"
)

const BLAST_SIZE int = 3

type Board struct {
	*engine.AbstractBoard
}

func (b *Board) GetHero() []engine.Point {
	return b.FindAllOf([]engine.Element{HERO, BOMB_HERO})
}

func (b *Board) GetOtherHeroes() []engine.Point {
	return b.FindAll(OTHER_HERO)
}

func (b *Board) IsMyHeroDead() bool {
	_, err := b.FindOne(DEAD_HERO)
	return err == nil
}

func (b *Board) IsBarrierAt(point engine.Point) bool {
	return b.IsAtAny(point, []engine.Element{
		HERO, BOMB_HERO, OTHER_HERO, OTHER_BOMB_HERO,
		BOMB_TIMER_5, BOMB_TIMER_4, BOMB_TIMER_3, BOMB_TIMER_2, BOMB_TIMER_1,
		WALL, DESTROYABLE_WALL, MEAT_CHOPPER,
	})
}

func (b *Board) GetBarriers() []engine.Point {
	return b.FindAllOf([]engine.Element{BOMB_HERO, OTHER_HERO, OTHER_BOMB_HERO, OTHER_DEAD_HERO,
		BOMB_TIMER_5, BOMB_TIMER_4, BOMB_TIMER_3, BOMB_TIMER_2, BOMB_TIMER_1, BOOM,
		WALL, DESTROYABLE_WALL, DESTROYED_WALL, MEAT_CHOPPER, DEAD_MEAT_CHOPPER})
}

func (b *Board) GetMeatChoppers() []engine.Point {
	return b.FindAll(MEAT_CHOPPER)
}

func (b *Board) GetWalls() []engine.Point {
	return b.FindAll(WALL)
}

func (b *Board) GetDestroyableWalls() []engine.Point {
	return b.FindAll(DESTROYABLE_WALL)
}

func (b *Board) GetBombs() []engine.Point {
	return b.FindAllOf([]engine.Element{BOMB_HERO, OTHER_BOMB_HERO,
		BOMB_TIMER_5, BOMB_TIMER_4, BOMB_TIMER_3, BOMB_TIMER_2, BOMB_TIMER_1})
}

func (b *Board) GetBlasts() []engine.Point {
	return b.FindAll(BOOM)
}

func (b *Board) GetPerks() []engine.Point {
	return b.FindAllOf([]engine.Element{BOMB_BLAST_RADIUS_INCREASE, BOMB_COUNT_INCREASE,
		BOMB_IMMUNE, BOMB_REMOTE_CONTROL})
}

func (b *Board) GetFutureBlasts() []engine.Point {
	var futureBlasts []engine.Point

	for _, bomb := range b.GetBombs() {
		// right
		for i := 1; i <= BLAST_SIZE; i++ {
			fBlast := engine.Point{X: bomb.X + i, Y: bomb.Y}
			if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
				break
			}
			futureBlasts = append(futureBlasts, fBlast)
		}
		// left
		for i := 1; i <= BLAST_SIZE; i++ {
			fBlast := engine.Point{X: bomb.X - i, Y: bomb.Y}
			if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
				break
			}
			futureBlasts = append(futureBlasts, fBlast)
		}
		// up
		for i := 1; i <= BLAST_SIZE; i++ {
			fBlast := engine.Point{X: bomb.X, Y: bomb.Y + i}
			if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
				break
			}
			futureBlasts = append(futureBlasts, fBlast)
		}
		// down
		for i := 1; i <= BLAST_SIZE; i++ {
			fBlast := engine.Point{X: bomb.X, Y: bomb.Y - i}
			if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
				break
			}
			futureBlasts = append(futureBlasts, fBlast)
		}
	}

	return removeDuplicates(futureBlasts)
}

func removeDuplicates(points []engine.Point) []engine.Point {
	set := make(map[engine.Point]struct{})
	for _, p := range points {
		if _, ok := set[p]; !ok {
			set[p] = struct{}{}
		}
	}
	res := make([]engine.Point, 0, len(set))
	for p, _ := range set {
		res = append(res, p)
	}
	return res
}
