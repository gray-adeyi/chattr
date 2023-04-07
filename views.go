package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

var rooms []Room
var users []User
var websocketConnections = []*websocket.Conn(nil)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}

func EnterChatRoom(c echo.Context) error {
	user := getOrCreateUser(c.FormValue("username"))
	roomName := c.FormValue("roomName")

	for _, room := range rooms {
		if room.RoomName == roomName {
			room.AddUser(&user)
			data := map[string]interface{}{
				"room":        room,
				"currentUser": user,
			}
			return c.Render(http.StatusOK, "room.html", data)
		}
	}
	newRoom := Room{
		Id:       uuid.New(),
		RoomName: roomName,
		Users:    []*User{&user},
	}
	rooms = append(rooms, newRoom)
	data := map[string]interface{}{
		"room":        newRoom,
		"currentUser": user,
	}
	return c.Render(http.StatusOK, "room.html", data)
}

func getOrCreateUser(username string) User {
	for _, user := range users {
		if user.Username == username {
			return user
		}
	}
	newUser := User{
		Id:       uuid.New(),
		Username: username,
	}
	users = append(users, newUser)
	return newUser
}

func ChatHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		websocketConnections = append(websocketConnections, ws)

		for {
			// Read
			msg := ""
			if err := websocket.Message.Receive(ws, &msg); err != nil && err != io.EOF {
				c.Logger().Error(err)
			}

			if msg == "" {
				continue
			}

			for _, conn := range websocketConnections {
				// Write
				if err := websocket.Message.Send(conn, msg); err != nil {
					c.Logger().Error(err)
				}
			}
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
