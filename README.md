## Mandatory steps
- copy and paste your game token (browserUrl)
```go
package main

...

func main() {
    // *** paste here board page url from a browser after registration ***
    browserUrl := "http://localhost:8080/codenjoy-contest/board/player/0?code=000000000000"
    ...
}
```

- make sure you've defined a correct game type of board & solver
```go
package main

import (
    ...
    "github.com/codenjoyme/codenjoy-go-client/games/bomberman"
    ...
)

func main() {
    ... 
    board := &bomberman.Board{AbstractBoard : &engine.AbstractBoard{}}
    solver := &bomberman.Solver{}
    ...
}
```

- implement your invincible solver algorithm
```go
package bomberman

...

type Solver struct {
}

func (s Solver) Get(b *Board) Action {
    ...
    // make your action
    return MoveFire(engine.UP)
}
...
```
