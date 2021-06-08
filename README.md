## Mandatory steps

- copy and paste your game token (browserUrl)
```go
package main

const URL = "http://localhost:8080/codenjoy-contest/board/player/0?code=000000000000"
```

- define the desirable type of game
```go
package main

const GAME = "bomberman"
```

- implement your invincible solver algorithm
```go
package bomberman

type Solver struct {
    B *Board
}

...

func (s Solver) nextStep() Action {
    // make your action
    return MoveFire(engine.UP)
}
```
