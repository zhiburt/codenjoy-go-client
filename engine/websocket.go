package engine

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"regexp"
	"strings"
)

const WebSocketContext = "/codenjoy-contest/ws"

func CreateWebSocketConnection(browserUrl string) (*websocket.Conn, error) {
	u, err := createWebUrl(browserUrl)
	if err != nil {
		return &websocket.Conn{}, err
	}

	log.Printf("Trying to esteblish connection with %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func createWebUrl(browserUrl string) (url.URL, error) {
	originalUrl := browserUrl

	var schema string
	if strings.Contains(browserUrl, "https://") {
		schema = "wss"
		browserUrl = strings.Replace(browserUrl, "https://", "", 1)
	} else if strings.Contains(browserUrl, "http://") {
		schema = "ws"
		browserUrl = strings.Replace(browserUrl, "http://", "", 1)
	} else {
		return url.URL{}, errors.New("Invalid URL. Unable to parse schema from `" + originalUrl + "`")
	}

	urlChunks := strings.Split(browserUrl, "/")
	if len(urlChunks) < 5 {
		return url.URL{}, errors.New("Invalid URL. Unable to parse host from `" + originalUrl + "`")
	}
	host := urlChunks[0]

	browserUrl = strings.Replace(browserUrl, host+"/codenjoy-contest/board/player/", "", 1)
	urlChunks = strings.Split(browserUrl, "?")
	r, e := containsNonWordCharacters(urlChunks[0])
	if r || e != nil {
		return url.URL{}, errors.New("Invalid URL. Unable to parse playerId from `" + originalUrl + "`")
	}
	playerId := urlChunks[0]

	browserUrl = strings.Replace(browserUrl, playerId+"?code=", "", 1)
	urlChunks = strings.Split(browserUrl, "&")
	r, e = containsNonDigitCharacters(urlChunks[0])
	if r || e != nil {
		return url.URL{}, errors.New("Invalid URL. Unable to parse code from `" + originalUrl + "`")
	}
	code := urlChunks[0]

	u := url.URL{
		Scheme:   schema,
		Host:     host,
		Path:     WebSocketContext,
		RawQuery: fmt.Sprintf("user=%s&code=%s", playerId, code),
	}
	return u, nil
}

func containsNonWordCharacters(input string) (bool, error) {
	return regexp.MatchString("\\W", input)
}

func containsNonDigitCharacters(input string) (bool, error) {
	return regexp.MatchString("\\D", input)
}
