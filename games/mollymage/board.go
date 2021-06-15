package mollymage

import (
    "github.com/codenjoyme/codenjoy-go-client/engine"
)

const BLAST_SIZE int = 3

type Board struct {
    *engine.AbstractBoard
}

func (b *Board) GetHero() []engine.Point {
    return b.FindAllOf([]engine.Element{HERO, POTION_HERO})
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
        HERO, POTION_HERO, OTHER_HERO, OTHER_POTION_HERO,
        POTION_TIMER_5, POTION_TIMER_4, POTION_TIMER_3, POTION_TIMER_2, POTION_TIMER_1,
        WALL, TREASURE_BOX, GHOST,
    })
}

func (b *Board) GetBarriers() []engine.Point {
    return b.FindAllOf([]engine.Element{POTION_HERO, OTHER_HERO, OTHER_POTION_HERO, OTHER_DEAD_HERO,
        POTION_TIMER_5, POTION_TIMER_4, POTION_TIMER_3, POTION_TIMER_2, POTION_TIMER_1, BOOM,
        WALL, TREASURE_BOX, OPENING_TREASURE_BOX, GHOST, DEAD_GHOST})
}

func (b *Board) GetGhosts() []engine.Point {
    return b.FindAll(GHOST)
}

func (b *Board) GetWalls() []engine.Point {
    return b.FindAll(WALL)
}

func (b *Board) GetTreasureBoxes() []engine.Point {
    return b.FindAll(TREASURE_BOX)
}

func (b *Board) GetPotions() []engine.Point {
    return b.FindAllOf([]engine.Element{POTION_HERO, OTHER_POTION_HERO,
        POTION_TIMER_5, POTION_TIMER_4, POTION_TIMER_3, POTION_TIMER_2, POTION_TIMER_1})
}

func (b *Board) GetBlasts() []engine.Point {
    return b.FindAll(BOOM)
}

func (b *Board) GetPerks() []engine.Point {
    return b.FindAllOf([]engine.Element{POTION_BLAST_RADIUS_INCREASE, POTION_COUNT_INCREASE,
        POTION_IMMUNE, POTION_REMOTE_CONTROL})
}

func (b *Board) GetFutureBlasts() []engine.Point {
    var futureBlasts []engine.Point

    for _, potion := range b.GetPotions() {
        // right
        for i := 1; i <= BLAST_SIZE; i++ {
            fBlast := engine.Point{X: potion.X + i, Y: potion.Y}
            if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
                break
            }
            futureBlasts = append(futureBlasts, fBlast)
        }
        // left
        for i := 1; i <= BLAST_SIZE; i++ {
            fBlast := engine.Point{X: potion.X - i, Y: potion.Y}
            if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
                break
            }
            futureBlasts = append(futureBlasts, fBlast)
        }
        // up
        for i := 1; i <= BLAST_SIZE; i++ {
            fBlast := engine.Point{X: potion.X, Y: potion.Y + i}
            if !b.IsValid(fBlast) || b.IsBarrierAt(fBlast) {
                break
            }
            futureBlasts = append(futureBlasts, fBlast)
        }
        // down
        for i := 1; i <= BLAST_SIZE; i++ {
            fBlast := engine.Point{X: potion.X, Y: potion.Y - i}
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
