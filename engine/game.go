package engine

import (
    "github.com/gorilla/websocket"
    "log"
    "strings"
)

type Communication struct {
    Done  chan struct{}
    Read  chan struct{}
    Write chan struct{}
}

type Envelope struct {
    Input  []rune
    Output string
}

func ReadWriteSocketLoop(connection *websocket.Conn, communication Communication, envelope *Envelope) {
    for {
        select {
        case <-communication.Done:
            log.Println("Client closed connection")
            return
        default:
            _, message, err := connection.ReadMessage()
            if err != nil {
                log.Println("Server closed connection, err:", err)
                close(communication.Done)
                return
            }

            rawBoard := strings.Replace(string(message), "board=", "", 1)
            envelope.Input = []rune(rawBoard)
            communication.Read <- struct{}{} // Make a signal to the client that it's time to make a move

            <-communication.Write // Wait for client response
            err = connection.WriteMessage(websocket.TextMessage, []byte(envelope.Output))
            if err != nil {
                log.Println("Failed to write command to the game server, err: ", err)
                close(communication.Done)
                return
            }
        }
    }
}
