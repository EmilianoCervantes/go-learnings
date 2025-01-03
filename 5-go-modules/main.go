package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const idExample = 12345

func saveUser(c echo.Context) error {
	returnMessage := fmt.Sprintf("Saved with id: %s", strconv.Itoa(idExample))
	return c.String(http.StatusOK, returnMessage)
}

// Example: http://localhost:1323/users/joe?team=x-men&member=wolverine
func getUser(c echo.Context) error {
	id := c.Param("id")
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	returnMessage := fmt.Sprintf("Here is the user: %s from the team %s and member name %s.", id, team, member)
	return c.String(http.StatusOK, returnMessage)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	returnMessage := "User with id:" + id
	return c.String(http.StatusOK, returnMessage)
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "User deleted")
}

func main() {
	myEcho := echo.New()

	myEcho.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World from Echo!")
	})

	// Routing examples
	myEcho.POST("/users", saveUser)
	myEcho.GET("/users/:id", getUser)
	myEcho.PUT("/users/:id", updateUser)
	myEcho.DELETE("/users/:id", deleteUser)

	myEcho.Logger.Fatal(myEcho.Start(":1323")) // http://127.0.0.1:1323/
}
