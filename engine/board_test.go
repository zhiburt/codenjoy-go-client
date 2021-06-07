package engine

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindOne(t *testing.T) {
	type tstruct struct {
		board          string
		element        Element
		expectedOutput Point
		expectedError  error
	}

	tests := []tstruct{
		{
			board: "☺.." +
				"..." +
				"...",
			element:        Element('☺'),
			expectedOutput: Point{0, 2},
			expectedError:  nil,
		},
		{
			board: "..☺" +
				"..." +
				"...",
			element:        Element('☺'),
			expectedOutput: Point{2, 2},
			expectedError:  nil,
		},
		{
			board: "..." +
				"..." +
				".☺.",
			element:        Element('☺'),
			expectedOutput: Point{1, 0},
			expectedError:  nil,
		},
		{
			board: "..." +
				"☺.." +
				"...",
			element:        Element('☺'),
			expectedOutput: Point{0, 1},
			expectedError:  nil,
		},
		{
			board: "..." +
				"..." +
				"...",
			element:        Element('☺'),
			expectedOutput: Point{-1, -1},
			expectedError:  errors.New("no such element"),
		},
	}

	b := AbstractBoard{}
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
		element        Element
		expectedOutput []Point
	}

	tests := []tstruct{
		{
			board: ".☺." +
				".☺☺" +
				"...",
			element:        Element('☺'),
			expectedOutput: []Point{{1, 2}, {1, 1}, {2, 1}},
		},
		{
			board: "..." +
				"..." +
				"...",
			element:        Element('☺'),
			expectedOutput: []Point(nil),
		},
	}

	b := AbstractBoard{}
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
		elements       []Element
		expectedOutput []Point
	}

	tests := []tstruct{
		{
			board: "..☺" +
				"Ѡ.." +
				".☻☻",
			elements:       []Element{'☺', 'Ѡ', '☻'},
			expectedOutput: []Point{{2, 2}, {0, 1}, {1, 0}, {2, 0}},
		},
		{
			board: "..." +
				"..." +
				"...",
			elements:       []Element{'☺', '☻', 'Ѡ'},
			expectedOutput: []Point(nil),
		},
	}

	b := AbstractBoard{}
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
		element        Element
		point          Point
		expectedOutput bool
	}

	tests := []tstruct{
		{
			board: "..." +
				"..☺" +
				"...",
			element:        Element('☺'),
			point:          Point{2, 1},
			expectedOutput: true,
		},
		{
			board: "..." +
				".☺." +
				"...",
			element:        Element('☺'),
			point:          Point{2, 1},
			expectedOutput: false,
		},
	}

	b := AbstractBoard{}
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
		elements       []Element
		point          Point
		expectedOutput bool
	}

	tests := []tstruct{
		{
			board: "..." +
				".☻." +
				".☺.",
			elements:       []Element{'☺', '☻'},
			point:          Point{1, 0},
			expectedOutput: true,
		},
		{
			board: "..." +
				".☻." +
				".☺.",
			elements:       []Element{'☺', '☻'},
			point:          Point{1, 1},
			expectedOutput: true,
		},
		{
			board: "..." +
				".☻." +
				".☺.",
			elements:       []Element{'☺', '☻'},
			point:          Point{1, 2},
			expectedOutput: false,
		},
	}

	b := AbstractBoard{}
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
		point          Point
		expectedOutput Element
		expectedError  error
	}

	tests := []tstruct{
		{
			board: "..." +
				".☺." +
				"...",
			point:          Point{-5, 1},
			expectedOutput: Element(' '),
			expectedError:  errors.New("invalid x value: -5"),
		},
		{
			board: "..." +
				".☺." +
				"...",
			point:          Point{1, 20},
			expectedOutput: Element(' '),
			expectedError:  errors.New("invalid y value: 20"),
		},
		{
			board: "..." +
				".☺." +
				"...",
			point:          Point{1, 1},
			expectedOutput: Element('☺'),
			expectedError:  nil,
		},
	}

	b := AbstractBoard{}
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
		element        Element
		point          Point
		expectedOutput bool
	}

	tests := []tstruct{
		{
			board: "..." +
				"..☺" +
				"...",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: true,
		},
		{
			board: ".☺." +
				"..." +
				"...",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: true,
		},
		{
			board: "..." +
				"☺.." +
				"...",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: true,
		},
		{
			board: "..." +
				"..." +
				".☺.",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: true,
		},
		{
			board: "..." +
				"..." +
				"..☺",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: false,
		},
	}

	b := AbstractBoard{}
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
		element        Element
		point          Point
		expectedOutput int
	}

	tests := []tstruct{
		{
			board: ".☺." +
				"..☺" +
				".☺.",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: 3,
		},
		{
			board: "☺.☺" +
				"..." +
				"☺.☺",
			element:        Element('☺'),
			point:          Point{1, 1},
			expectedOutput: 0,
		},
	}

	b := AbstractBoard{}
	for _, tt := range tests {
		t.Run("count element near point", func(t *testing.T) {
			b.UpdateBoard([]rune(tt.board))
			output := b.CountNear(tt.point, tt.element)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
