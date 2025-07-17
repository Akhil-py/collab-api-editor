package handlers

import (
    "github.com/labstack/echo/v4"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func WSHandler(c echo.Context) error {
    // TODO: Validate JWT from query/header
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer conn.Close()
    // TODO: Handle messages, presence, etc.
    return nil
} 