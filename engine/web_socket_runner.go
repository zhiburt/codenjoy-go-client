package engine

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"regexp"
	"strings"
)

const URL_REGEX = "(?P<scheme>http|https)://(?P<host>.+)/codenjoy-contest/board/player/(?P<player>\\w+)\\?code=(?P<code>\\d+)"

type WebSocketRunner struct {
	token string
}

func NewWebSocketRunner(url string) *WebSocketRunner {
	return &WebSocketRunner{UrlToWsToken(url)}
}

func UrlToWsToken(url string) string {
	r := regexp.MustCompile(URL_REGEX)
	params := r.FindStringSubmatch(url)

	scheme := params[r.SubexpIndex("scheme")]
	if strings.EqualFold(scheme, "https") {
		scheme = "wss"
	} else {
		scheme = "ws"
	}
	return fmt.Sprintf("%s://%s/codenjoy-contest/ws?user=%s&code=%s",
		scheme,
		params[r.SubexpIndex("host")],
		params[r.SubexpIndex("player")],
		params[r.SubexpIndex("code")])
}

func (runner *WebSocketRunner) Run(solver Solver) {
	connection, _, err := websocket.DefaultDialer.Dial(runner.token, nil)
	if err != nil {
		log.Println("unable to establish websocket connection, error: ", err)
		return
	}
	log.Println("websocket connection successfully established")

	for {
		_, msgFromServer, err := connection.ReadMessage()
		if err != nil {
			log.Println("websocket ReadMessage error: ", err)
			return
		}
		msgToServer := solver.Answer(string(msgFromServer))
		err = connection.WriteMessage(websocket.TextMessage, []byte(msgToServer))
		if err != nil {
			log.Println("websocket WriteMessage error: ", err)
			return
		}
	}
}
