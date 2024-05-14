package main

import (
	"fmt"
)

// Mediator design pattern
type Message interface {
	AddUser(user *User)
	SendMessage(sender *User, message string)
}

type ChatRoom struct {
	users []*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make([]*User, 0),
	}
}

func (cr *ChatRoom) AddUser(user *User) {
	cr.users = append(cr.users, user)
}

func (cr *ChatRoom) SendMessage(sender *User, message string) {
	for _, user := range cr.users {
		if user != sender {
			user.ReceiveMessage(sender, message)
		}
	}
}

type User struct {
	ID   int
	Name string
	Room Message
}

func NewUser(id int, name string, room Message) *User {
	return &User{
		ID:   id,
		Name: name,
		Room: room,
	}
}

func (u *User) Send(message string) {
	u.Room.SendMessage(u, message)
}

func (u *User) ReceiveMessage(sender *User, message string) {
	fmt.Printf("User %s received message from %s: %s\n", u.Name, sender.Name, message)
}

func main() {
	chatRoom := NewChatRoom()

	ram := NewUser(1, "Ram", chatRoom)
	jay := NewUser(2, "Jay", chatRoom)
	sanket := NewUser(3, "Sanket", chatRoom)

	chatRoom.AddUser(ram)
	chatRoom.AddUser(jay)
	chatRoom.AddUser(sanket)

	ram.Send("Hello everyone!")
	jay.Send("Hi Ram..!")
	sanket.Send("Hey there!")

}
