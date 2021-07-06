package engine

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyMessage(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "")
	assert.Equal(t, "", board.String())
}

func TestEmptySupportedElements(t *testing.T) {
	assert.Panics(t, func() { engine.NewGameBoard([]rune{}, "aaa"+"bbb"+"ccc") })
}

func TestValidMessageAndSupportedElements(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, "aaa\nbbb\nccc\n", board.String())
}

func TestEraseMessagePrefix(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "board="+"aaa"+"bbb"+"ccc")
	assert.Equal(t, "aaa\nbbb\nccc\n", board.String())
}

func TestMessageWithUnsupportedElements(t *testing.T) {
	assert.Panics(t, func() { engine.NewGameBoard([]rune{'a', 'b', 'c'}, "ab8c") })
}

func TestGetSize(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, 3, board.GetSize())
}

func TestGetAt(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, 'c', board.GetAt(engine.NewPoint(0, 0)))
	assert.Equal(t, 'c', board.GetAt(engine.NewPoint(1, 0)))
	assert.Equal(t, 'c', board.GetAt(engine.NewPoint(2, 0)))
	assert.Equal(t, 'b', board.GetAt(engine.NewPoint(0, 1)))
	assert.Equal(t, 'b', board.GetAt(engine.NewPoint(1, 1)))
	assert.Equal(t, 'b', board.GetAt(engine.NewPoint(2, 1)))
	assert.Equal(t, 'a', board.GetAt(engine.NewPoint(0, 2)))
	assert.Equal(t, 'a', board.GetAt(engine.NewPoint(1, 2)))
	assert.Equal(t, 'a', board.GetAt(engine.NewPoint(2, 2)))
}

func TestGetAtInvalidPoint(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Panics(t, func() { board.GetAt(engine.NewPoint(10, 10)) })
}

func TestFind(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, "[[0,2] [1,2] [2,2]]", fmt.Sprintf("%v",
		board.Find('a')))
	assert.Equal(t, "[[0,0] [0,1] [1,0] [1,1] [2,0] [2,1]]", fmt.Sprintf("%v",
		board.Find('b', 'c')))
}

func TestFindNotExistedElement(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, "[]", fmt.Sprintf("%v", board.Find('d')))
}

func TestFindFirst(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.FindFirst('c')))
	assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindFirst('b', 'c')))
	assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindFirst('c', 'b')))
}

func TestFindFirstNotExistedElement(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, "<nil>", fmt.Sprintf("%v", board.FindFirst('d')))
}

func TestIsAt(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, true, board.IsAt(engine.NewPoint(1, 2), 'a'))
	assert.Equal(t, false, board.IsAt(engine.NewPoint(1, 2), 'b'))
	assert.Equal(t, false, board.IsAt(engine.NewPoint(1, 2), 'c'))
}

func TestIsAtInvalidPoint(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
	assert.Equal(t, false, board.IsAt(engine.NewPoint(10, 10), 'b'))
}

func TestFindNear(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, []rune{'f', 'd', 'b', 'h'}, board.FindNear(engine.NewPoint(1, 1)))
}

func TestFindNearInvalidPoint(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, []rune(nil), board.FindNear(engine.NewPoint(-1, -1)))
}

func TestCountNear(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, 2, board.CountNear(engine.NewPoint(1, 1), 'a', 'b', 'c', 'd'))
}

func TestCountNearInvalidPoint(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, 0, board.CountNear(engine.NewPoint(-1, -1), 'a', 'b', 'c', 'd'))
}

func TestCountNearNotExistedElement(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, 0, board.CountNear(engine.NewPoint(1, 1), 'r'))
	assert.Equal(t, 0, board.CountNear(engine.NewPoint(1, 1), 'x', 'y', 'z'))
}

func TestIsNear(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, false, board.IsNear(engine.NewPoint(1, 1), 'a'))
	assert.Equal(t, true, board.IsNear(engine.NewPoint(1, 1), 'b'))
	assert.Equal(t, true, board.IsNear(engine.NewPoint(1, 1), 'c', 'd'))
}

func TestIsNearInvalidPoint(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, false, board.IsNear(engine.NewPoint(-1, -1), 'a'))
}

func TestIsNearNotExistedElement(t *testing.T) {
	board := engine.NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
	assert.Equal(t, false, board.IsNear(engine.NewPoint(1, 1), 'r'))
	assert.Equal(t, false, board.IsNear(engine.NewPoint(1, 1), 'x', 'y', 'z'))
}
