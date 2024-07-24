package tracking

import "github.com/labstack/echo/v4"

type HttpHandler interface {
	SaveTracking(c echo.Context) error
}
