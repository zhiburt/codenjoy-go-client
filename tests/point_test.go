package tests

import (
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_IndexToPoint(t *testing.T) {
    type tstruct struct {
        board          string
        index          int
        expectedOutput engine.Point
    }

    tests := []tstruct{
        {
            board:  ".☺." +
                    "☺.." +
                    "..☺",
            index:          1,
            expectedOutput: engine.Point{X: 1, Y: 2},
        },
        {
            board:  ".☺." +
                    "☺.." +
                    "..☺",
            index:          3,
            expectedOutput: engine.Point{X: 0, Y: 1},
        },
        {
            board:  ".☺." +
                    "☺.." +
                    "..☺",
            index:          8,
            expectedOutput: engine.Point{X: 2, Y: 0},
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("convert index to point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.IndexToPoint(tt.index)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_PointToIndex(t *testing.T) {
    type tstruct struct {
        board          string
        point          engine.Point
        expectedOutput int
    }

    tests := []tstruct{
        {
            board:  "☺.." +
                    "..☺" +
                    ".☺.",
            point:          engine.Point{X: 0, Y: 2},
            expectedOutput: 0,
        },
        {
            board:  "☺.." +
                    "..☺" +
                    ".☺.",
            point:          engine.Point{X: 2, Y: 1},
            expectedOutput: 5,
        },
        {
            board:  "☺.." +
                    "..☺" +
                    ".☺.",
            point:          engine.Point{X: 1, Y: 0},
            expectedOutput: 7,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("convert point to index", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.PointToIndex(tt.point)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}
