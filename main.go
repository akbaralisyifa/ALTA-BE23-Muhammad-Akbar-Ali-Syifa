package main

import (
	"todos/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // initial echo
	factory.InitalFactory(e);
	e.Start(":8000")
}