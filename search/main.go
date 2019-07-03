package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "yakushou730@gmail.com", Name: "Yao Shang Tseng"},
	{Email: "shou.tseng+1@honestbee.com", Name: "Shou Tseng 1"},
	{Email: "shou.tseng+2@honestbee.com", Name: "Shou Tseng 2"},
	{Email: "shou.tseng+3@honestbee.com", Name: "Shou Tseng 3"},
	{Email: "shou.tseng+4@honestbee.com", Name: "Shou Tseng 4"},
	{Email: "shou.tseng+5@honestbee.com", Name: "Shou Tseng 5"},
	{Email: "shou.tseng+6@honestbee.com", Name: "Shou Tseng 6"},
	{Email: "shou.tseng+7@honestbee.com", Name: "Shou Tseng 7"},
}

type Worker struct {
	users []User
	ch    chan *User
	name  string
}

func NewWorker(users []User, ch chan *User, name string) *Worker {
	return &Worker{users: users, ch: ch, name: name}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if strings.Contains(user.Email, email) {
			log.Println(">>", w.name)
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]

	ch := make(chan *User)

	log.Printf("Looking for %s", email)
	go NewWorker(DataBase[:3], ch, "#1").Find(email)
	go NewWorker(DataBase[3:6], ch, "#2").Find(email)
	go NewWorker(DataBase[6:], ch, "#3").Find(email)

	for {
		select {
		case user := <-ch:
			log.Printf("The email %s is owned by %s", user.Email, user.Name)
		case <-time.After(1 * time.Second):
			return
		}
	}
}
