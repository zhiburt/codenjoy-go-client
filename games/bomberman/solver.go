package bomberman

import (
    "fmt"
    "github.com/codenjoyme/codenjoy-go-client/engine"
)

type Solver struct {
}

func (s Solver) Get(b *Board) Action {
    printDetailedBoardInfo(b)
    // make your action
    return MoveFire(engine.UP)
}

func printDetailedBoardInfo(b *Board) {
    fmt.Printf("Hero at: %v\n", b.GetHero())
    fmt.Printf("Other heroes at: %v\n", b.GetOtherHeroes())
    fmt.Printf("Meat choppers at: %v\n", b.GetMeatChoppers())
    fmt.Printf("Destroy walls at: %v\n", b.GetDestroyableWalls())
    fmt.Printf("Bombs at: %v\n", b.GetBombs())
    fmt.Printf("Blasts at: %v\n", b.GetBlasts())
    fmt.Printf("Expected blasts at: %v\n", b.GetFutureBlasts())
}
