package mollymage

import (
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/codenjoyme/codenjoy-go-client/games/mollymage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Should_When(t *testing.T) {
	solver := mollymage.Solver{B: &mollymage.Board{AbstractBoard: &engine.AbstractBoard{}}}
	board := []rune("☼☼☼☼☼☼☼☼☼☼☼☼☼☼☼" +
					"☼☺        # # ☼" +
					"☼ ☼ ☼ ☼#☼ ☼ ☼ ☼" +
					"☼##           ☼" +
					"☼ ☼ ☼#☼ ☼ ☼ ☼ ☼" +
					"☼   #    # #  ☼" +
					"☼ ☼ ☼ ☼#☼ ☼ ☼ ☼" +
					"☼             ☼" +
					"☼#☼ ☼ ☼#☼ ☼ ☼#☼" +
					"☼  #  #       ☼" +
					"☼ ☼ ☼ ☼ ☼ ☼ ☼#☼" +
					"☼ ##      #   ☼" +
					"☼ ☼ ☼ ☼ ☼ ☼ ☼#☼" +
					"☼ #   #  &    ☼" +
					"☼☼☼☼☼☼☼☼☼☼☼☼☼☼☼")

	t.Run("get next step", func(t *testing.T) {
		expect := []rune("UP,ACT")
		actual := []rune(solver.Get(board))
		assert.ElementsMatch(t, expect, actual)
	})
}
