package main

import (
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/codenjoyme/codenjoy-go-client/games/mollymage"
	"log"
	"os"
)

const URL = "http://localhost:8080/codenjoy-contest/board/player/0?code=000000000000"
const GAME = "mollymage"

type Solver interface {
	Get(rawBoard []rune) string
}

func main() {
	communication, envelope, err := engine.CreateWebSocketConnection(loadBrowserUrl())
	if err != nil {
		log.Fatal(err)
		return
	}

	solver := loadGameSolver()
	for {
		select {
		case <-communication.Done:
			log.Fatal("It's done")
			return
		case <-communication.Read:
			envelope.Output = solver.Get(envelope.Input)
			communication.Write <- struct{}{}
		}
	}
}

func loadBrowserUrl() string {
	url := URL
	// os.Args[0] the name of command itself; os.Args[1] game; os.Args[2] url
	if len(os.Args) == 3 {
		url = os.Args[2]
	}
	return url
}

func loadGameSolver() Solver {
	game := GAME
	// os.Args[0] the name of command itself; os.Args[1] game; os.Args[2] url
	if len(os.Args) == 3 {
		game = os.Args[1]
	}

	switch game {
	case "mollymage":
		return &mollymage.Solver{B: &mollymage.Board{AbstractBoard: &engine.AbstractBoard{}}}
	}
	panic("unable to determine game type")
}
