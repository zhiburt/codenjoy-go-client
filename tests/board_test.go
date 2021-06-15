package tests

import (
    "errors"
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_FindOne(t *testing.T) {
    type tstruct struct {
        board          string
        element        engine.Element
        expectedOutput engine.Point
        expectedError  error
    }

    tests := []tstruct{
        {
            board:  "☺.." +
                    "..." +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: engine.Point{X: 0, Y: 2},
            expectedError:  nil,
        },
        {
            board:  "..☺" +
                    "..." +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: engine.Point{X: 2, Y: 2},
            expectedError:  nil,
        },
        {
            board:  "..." +
                    "..." +
                    ".☺.",
            element:        engine.Element('☺'),
            expectedOutput: engine.Point{X: 1, Y: 0},
            expectedError:  nil,
        },
        {
            board:  "..." +
                    "☺.." +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: engine.Point{X: 0, Y: 1},
            expectedError:  nil,
        },
        {
            board:  "..." +
                    "..." +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: engine.Point{X: -1, Y: -1},
            expectedError:  errors.New("no such element"),
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("find element point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output, err := b.FindOne(tt.element)
            assert.Equal(t, tt.expectedOutput, output)
            assert.Equal(t, tt.expectedError, err)
        })
    }
}

func Test_FindAll(t *testing.T) {
    type tstruct struct {
        board          string
        element        engine.Element
        expectedOutput []engine.Point
    }

    tests := []tstruct{
        {
            board:  ".☺." +
                    ".☺☺" +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: []engine.Point{{1, 2}, {1, 1}, {2, 1}},
        },
        {
            board:  "..." +
                    "..." +
                    "...",
            element:        engine.Element('☺'),
            expectedOutput: []engine.Point(nil),
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("find all element points", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.FindAll(tt.element)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_FindAllOf(t *testing.T) {
    type tstruct struct {
        board          string
        elements       []engine.Element
        expectedOutput []engine.Point
    }

    tests := []tstruct{
        {
            board:  "..☺" +
                    "Ѡ.." +
                    ".☻☻",
            elements:       []engine.Element{'☺', 'Ѡ', '☻'},
            expectedOutput: []engine.Point{{2, 2}, {0, 1}, {1, 0}, {2, 0}},
        },
        {
            board:  "..." +
                    "..." +
                    "...",
            elements:       []engine.Element{'☺', '☻', 'Ѡ'},
            expectedOutput: []engine.Point(nil),
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("find all elements points", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.FindAllOf(tt.elements)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_IsAt(t *testing.T) {
    type tstruct struct {
        board          string
        element        engine.Element
        point          engine.Point
        expectedOutput bool
    }

    tests := []tstruct{
        {
            board:  "..." +
                    "..☺" +
                    "...",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 2, Y: 1},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    ".☺." +
                    "...",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 2, Y: 1},
            expectedOutput: false,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("is element at point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.IsAt(tt.point, tt.element)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_IsAtAny(t *testing.T) {
    type tstruct struct {
        board          string
        elements       []engine.Element
        point          engine.Point
        expectedOutput bool
    }

    tests := []tstruct{
        {
            board:  "..." +
                    ".☻." +
                    ".☺.",
            elements:       []engine.Element{'☺', '☻'},
            point:          engine.Point{X: 1, Y: 0},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    ".☻." +
                    ".☺.",
            elements:       []engine.Element{'☺', '☻'},
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    ".☻." +
                    ".☺.",
            elements:       []engine.Element{'☺', '☻'},
            point:          engine.Point{X: 1, Y: 2},
            expectedOutput: false,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("is any element at point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.IsAtAny(tt.point, tt.elements)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_GetAt(t *testing.T) {
    type tstruct struct {
        board          string
        point          engine.Point
        expectedOutput engine.Element
        expectedError  error
    }

    tests := []tstruct{
        {
            board:  "..." +
                    ".☺." +
                    "...",
            point:          engine.Point{X: -5, Y: 1},
            expectedOutput: engine.Element(' '),
            expectedError:  errors.New("invalid x value: -5"),
        },
        {
            board:  "..." +
                    ".☺." +
                    "...",
            point:          engine.Point{X: 1, Y: 20},
            expectedOutput: engine.Element(' '),
            expectedError:  errors.New("invalid y value: 20"),
        },
        {
            board:  "..." +
                    ".☺." +
                    "...",
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: engine.Element('☺'),
            expectedError:  nil,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("get element at point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output, err := b.GetAt(tt.point)
            assert.Equal(t, tt.expectedOutput, output)
            assert.Equal(t, tt.expectedError, err)
        })
    }
}

func Test_IsNear(t *testing.T) {
    type tstruct struct {
        board          string
        element        engine.Element
        point          engine.Point
        expectedOutput bool
    }

    tests := []tstruct{
        {
            board:  "..." +
                    "..☺" +
                    "...",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: true,
        },
        {
            board:  ".☺." +
                    "..." +
                    "...",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    "☺.." +
                    "...",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    "..." +
                    ".☺.",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: true,
        },
        {
            board:  "..." +
                    "..." +
                    "..☺",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: false,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("is element near point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.IsNear(tt.point, tt.element)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}

func Test_CountNear(t *testing.T) {
    type tstruct struct {
        board          string
        element        engine.Element
        point          engine.Point
        expectedOutput int
    }

    tests := []tstruct{
        {
            board:  ".☺." +
                    "..☺" +
                    ".☺.",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: 3,
        },
        {
            board:  "☺.☺" +
                    "..." +
                    "☺.☺",
            element:        engine.Element('☺'),
            point:          engine.Point{X: 1, Y: 1},
            expectedOutput: 0,
        },
    }

    b := engine.AbstractBoard{}
    for _, tt := range tests {
        t.Run("count element near point", func(t *testing.T) {
            b.UpdateBoard([]rune(tt.board))
            output := b.CountNear(tt.point, tt.element)
            assert.Equal(t, tt.expectedOutput, output)
        })
    }
}
