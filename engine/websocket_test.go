package engine

import (
    "errors"
    "github.com/stretchr/testify/assert"
    "net/url"
    "testing"
)

func Test_createWebUrl(t *testing.T) {
    type tstruct struct {
        title          string
        input          string
        expectedWebUrl url.URL
        expectedError  error
    }

    tests := []tstruct{
        {
            title: "Valid http URL",
            input: "http://127.0.0.1:8080/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            expectedWebUrl: url.URL{
                Scheme:   "ws",
                Host:     "127.0.0.1:8080",
                Path:     WebSocketContext,
                RawQuery: "user=793wdxskw521spo4mn1y&code=531459153668826800",
            },
            expectedError: nil,
        },
        {
            title: "Valid https URL",
            input: "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            expectedWebUrl: url.URL{
                Scheme:   "wss",
                Host:     "dojorena.io",
                Path:     WebSocketContext,
                RawQuery: "user=793wdxskw521spo4mn1y&code=531459153668826800",
            },
            expectedError: nil,
        },
        {
            title:          "Unsupported schema",
            input:          "ftp://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            expectedWebUrl: url.URL{},
            expectedError:  errors.New("Invalid URL. Unable to parse schema from `ftp://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800`"),
        },
        {
            title:          "Invalid host",
            input:          "https://codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            expectedWebUrl: url.URL{},
            expectedError:  errors.New("Invalid URL. Unable to parse host from `https://codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800`"),
        },
        {
            title:          "PlayerId contains non word characters",
            input:          "https://dojorena.io/codenjoy-contest/board/player/7**wdxskw521spo4mn1y?code=531459153668826800",
            expectedWebUrl: url.URL{},
            expectedError:  errors.New("Invalid URL. Unable to parse playerId from `https://dojorena.io/codenjoy-contest/board/player/7**wdxskw521spo4mn1y?code=531459153668826800`"),
        },
        {
            title:          "Code contains non digit characters",
            input:          "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=5AA459153668826800",
            expectedWebUrl: url.URL{},
            expectedError:  errors.New("Invalid URL. Unable to parse code from `https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=5AA459153668826800`"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.title, func(t *testing.T) {
            webUrl, err := createWebUrl(tt.input)
            assert.Equal(t, tt.expectedWebUrl, webUrl)
            assert.Equal(t, tt.expectedError, err)
        })
    }
}
