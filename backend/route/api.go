package route

import (
	"github.com/Blackmocca/example-geolocation/middleware"
	"github.com/Blackmocca/example-geolocation/service/tracking"
	_track_validator "github.com/Blackmocca/example-geolocation/service/tracking/validator"
	"github.com/labstack/echo/v4"
)

type Route struct {
	e     *echo.Echo
	middl middleware.GoMiddlewareInf
}

func NewRoute(e *echo.Echo, middl middleware.GoMiddlewareInf) *Route {
	return &Route{e: e, middl: middl}
}

func (r *Route) Register(handler tracking.HttpHandler, validator _track_validator.Validation) {
	r.e.POST("/trackings", handler.SaveTracking, validator.ValidateSaveTracking)
}
