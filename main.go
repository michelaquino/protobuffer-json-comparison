package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/michelaquino/protobuffer-json-comparison/handlers"
)

const (
	port = 8888
)

func main() {
	echoInstance := echo.New()
	echoInstance.GET("/json", handlers.GetJSON)
	echoInstance.POST("/json", handlers.PostJSON)
	echoInstance.GET("/proto", handlers.GetPROTO)
	echoInstance.POST("/proto", handlers.PostPROTO)
	echoInstance.GET("/protojson", handlers.GetPROTOAsJSON)
	echoInstance.POST("/protojson", handlers.PostPROTOAsJSON)

	echoInstance.GET("/comparison", handlers.PrintComparison)

	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", port)))
}
