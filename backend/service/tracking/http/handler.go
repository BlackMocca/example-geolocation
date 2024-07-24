package http

import (
	"net/http"

	"github.com/Blackmocca/example-geolocation/models"
	"github.com/Blackmocca/example-geolocation/service/tracking"
	"github.com/Blackmocca/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type trackingHandler struct {
	trackingUs tracking.Usecase
}

func NewTrackingHandler(trackingUs tracking.Usecase) tracking.HttpHandler {
	return &trackingHandler{trackingUs: trackingUs}
}

func (t trackingHandler) SaveTracking(c echo.Context) error {
	var ctx = c.Request().Context()
	var params = c.Get("params").(map[string]interface{})

	var tracking = &models.Tracking{
		PlateNumber: cast.ToString(params["plate_number"]),
		Lat:         cast.ToFloat64(params["lat"]),
		Lon:         cast.ToFloat64(params["lon"]),
		Tracktime:   utils.NewTimestampFromNow().ToPointer(),
	}
	if v, ok := params["track_time"]; ok {
		tracking.Tracktime = utils.NewTimestampFromString(cast.ToString(v)).ToPointer()
	}

	if err := t.trackingUs.SaveTracking(ctx, tracking); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"message": "successful",
	}
	return c.JSON(http.StatusOK, resp)
}
