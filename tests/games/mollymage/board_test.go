package mollymage

import (
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/codenjoyme/codenjoy-go-client/games/mollymage"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_GetFutureBlasts(t *testing.T) {
    type tstruct struct {
        board          string
        expectedOutput []engine.Point
    }

    tests := []tstruct{
        {
            board:  ".☼☼.." +
                    ".3☼.." +
                    ".☼..." +
                    "..&2&" +
                    "1♥...",
            expectedOutput: []engine.Point{{0, 3}, {3, 2}, {3, 3}, {3, 4}, {3, 0}, {0, 1}, {0, 2}},
        },
    }

    b := mollymage.Board{AbstractBoard: &engine.AbstractBoard{}}
    for _, tt := range tests {
        t.Run("get future blasts", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.GetFutureBlasts()
            assert.ElementsMatch(t, tt.expectedOutput, output)
        })
    }
}
