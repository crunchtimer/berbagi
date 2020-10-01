package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {

	ctrls, err := NewControllers()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		err := ctrls.App.TallyCount()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
