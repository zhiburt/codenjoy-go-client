package mollymage

import "github.com/codenjoyme/codenjoy-go-client/engine"

type Action string

const (
    ACT  = "ACT"
    STOP = "STOP"

    NOTHING = ""
)

func (a Action) IsValid() bool {
    switch a {
    case "UP", "DOWN", "RIGHT", "LEFT", "ACT", "STOP", "UP,ACT", "DOWN,ACT", "RIGHT,ACT", "LEFT,ACT", "ACT,UP", "ACT,DOWN", "ACT,RIGHT", "ACT,LEFT":
        return true
    default:
        return false
    }
}

func Act() Action {
    return ACT
}

func Move(d engine.Direction) Action {
    return Action(d.String())
}

func MoveFire(d engine.Direction) Action {
    return Action(d.String() + "," + ACT)
}

func FireMove(d engine.Direction) Action {
    return Action(ACT + "," + d.String())
}

func Stop() Action {
    return STOP
}
