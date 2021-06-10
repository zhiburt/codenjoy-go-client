package engine

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_IndexToPoint(t *testing.T) {
    type tstruct struct {
        board          string
        index          int
        expectedOutput Point
    }

    tests := []tstruct{
        {
            board: ".☺." +
                "☺.." +
                "..☺",
            index:          1,
            expectedOutput: Point{1, 2},
        },
        {
            board: ".☺." +
                "☺.." +
                "..☺",
            index:          3,
            expectedOutput: Point{0, 1},
        },
        {
            board: ".☺." +
                "☺.." +
                "..☺",
            index:          8,
            expectedOutput: Point{2, 0},
        },
    }

    b := AbstractBoard{}
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
        point          Point
        expectedOutput int
    }

    tests := []tstruct{
        {
            board: "☺.." +
                "..☺" +
                ".☺.",
            point:          Point{0, 2},
            expectedOutput: 0,
        },
        {
            board: "☺.." +
                "..☺" +
                ".☺.",
            point:          Point{2, 1},
            expectedOutput: 5,
        },
        {
            board: "☺.." +
                "..☺" +
                ".☺.",
            point:          Point{1, 0},
            expectedOutput: 7,
        },
    }

    b := AbstractBoard{}
    for _, tt := range tests {
        t.Run("convert point to index", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.PointToIndex(tt.point)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}
