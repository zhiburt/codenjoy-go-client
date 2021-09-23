package clifford

import (
	"github.com/codenjoyme/codenjoy-go-client/games/clifford"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnswer(t *testing.T) {
	message := "board=" +
		"☼☼☼☼☼" +
		"☼   ☼" +
		"☼ ⊳ ☼" +
		"☼   ☼" +
		"☼☼☼☼☼"
	assert.Equal(t, "ACT", clifford.NewSolver().Answer(message))
}
