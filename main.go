package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michelaquino/protobuffer-json-comparison/handlers"
)

const (
	port = 8888
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/json", handlers.GetJSON)
	e.POST("/json", handlers.PostJSON)

	e.GET("/proto", handlers.GetPROTO)
	e.POST("/proto", handlers.PostPROTO)

	e.GET("/protojson", handlers.GetPROTOAsJSON)
	e.POST("/protojson", handlers.PostPROTOAsJSON)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
