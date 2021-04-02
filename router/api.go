package router

import (
	"github.com/labstack/echo/v4"
	"home-test-tiki/handler"
)

type API struct {
	Echo *echo.Echo
	StringHandler handler.StringHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)
	api.Echo.GET("/search/", api.StringHandler.HandlerSplitString)
}