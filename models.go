package main

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type Room struct {
	Id       uuid.UUID `json:"id"`
	RoomName string    `json:"roomName"`
	Users    []*User   `json:"users"`
}

func (r *Room) AddUser(user *User) {
	for _, userInRoom := range r.Users {
		if userInRoom.Id == user.Id {
			return
		}
	}
	r.Users = append(r.Users, user)
	return
}
