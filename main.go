package main

import "github.com/codenjoyme/codenjoy-go-client/engine"

func main() {
	// *** paste here board page url from a browser after registration ***
	engine.CreateWebSocketConnection("http://localhost:8080/codenjoy-contest/board/player/0?code=000000000000")
}
