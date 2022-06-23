package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Chào mừng bạn đã đến Home Test Tiki được viết bởi Nimtechnology, Query path /search/?message=string(replace message).")
}
