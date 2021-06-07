package main

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/codenjoyme/codenjoy-go-client/games/bomberman"
	"log"
)

func main() {
	// *** paste here board page url from a browser after registration ***
	browserUrl := "http://localhost:8080/codenjoy-contest/board/player/0?code=000000000000"
	communication, envelope, err := engine.CreateWebSocketConnection(browserUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	board := &bomberman.Board{AbstractBoard: &engine.AbstractBoard{}}
	solver := &bomberman.Solver{}

	for {
		select {
		case <-communication.Done:
			log.Fatal("It's done")
			return
		case <-communication.Read:
			board.UpdateBoard(envelope.Input)
			fmt.Printf("\nBoard:\n%s\n", board.BoardAsString())

			envelope.Output = string(solver.Get(board))
			fmt.Println("Answer: " + envelope.Output)
			fmt.Println("-------------------------------------------------------------")

			communication.Write <- struct{}{}
		}
	}
}
