package main

import (
	"MyLink_Server/server/internal/app"
	netcompete "MyLink_Server/server/internal/app/handler/competition"
)

func init() {
	netcompete.Rooms = make(map[int]*netcompete.Room)
}

func main() {
	serv := app.NewServer()
	serv.Run()
}
