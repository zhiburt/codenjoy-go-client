package engine

import (
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestValidUrlToWstoken(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wstoken string
    }{
        {
            name:    "UrlToWsToken with valid HTTP URL",
            url:     "http://127.0.0.1:8080/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            wstoken: "ws://127.0.0.1:8080/codenjoy-contest/ws?user=793wdxskw521spo4mn1y&code=531459153668826800",
        },
        {
            name:    "UrlToWsToken with valid HTTPS URL",
            url:     "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            wstoken: "wss://dojorena.io/codenjoy-contest/ws?user=793wdxskw521spo4mn1y&code=531459153668826800",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := engine.UrlToWsToken(tt.url); got != tt.wstoken {
                t.Errorf("urlToWsToken() = %v, want %v", got, tt.wstoken)
            }
        })
    }
}

func Test_Panic_UrlToWsToken(t *testing.T) {
    tests := []struct {
        name string
        url  string
    }{
        {
            name: "UrlToWsToken with unsupported scheme",
            url:  "ftp://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "UrlToWsToken with invalid host",
            url:  "https://codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "UrlToWsToken with playerId that contains non word characters",
            url:  "https://dojorena.io/codenjoy-contest/board/player/7**wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "UrlToWsToken with code that contains non word characters",
            url:  "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=AA5459153668826800",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Panics(t, func() { engine.UrlToWsToken(tt.url) })
        })
    }
}
