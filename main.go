package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Player struct {
	Id       int    `json: "id"`
	Nickname string `json: "nickname"`
	Online   bool   `json: "online"`
}

type Players []Player

var players Player

func getPlayers(c echo.Context) error {
	return c.JSON(http.StatusOK, players)
}

func main() {
	fmt.Println("ahhhhhhhhhh")
	e := echo.New()
	e.Get("/players", getPlayers)
	e.Start(":1256")
}
