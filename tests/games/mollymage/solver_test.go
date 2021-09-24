package mollymage

import (
    "github.com/codenjoyme/codenjoy-go-client/games/mollymage"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAnswer(t *testing.T) {
    message := "board=" +
        "☼☼☼☼☼" +
        "☼   ☼" +
        "☼ ☺ ☼" +
        "☼   ☼" +
        "☼☼☼☼☼"
    assert.Equal(t, "ACT", mollymage.NewSolver().Answer(message))
}
