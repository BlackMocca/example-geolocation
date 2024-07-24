package validator

import (
	"net/http"

	"github.com/Blackmocca/example-geolocation/middleware"
	"github.com/labstack/echo/v4"
	"github.com/xeipuuv/gojsonschema"
)

type Validation struct {
	SaveTrackingSchema *gojsonschema.Schema
}

func (v Validation) ValidateSaveTracking(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params = c.Get("params").(map[string]interface{})

		result, err := v.SaveTrackingSchema.Validate(gojsonschema.NewGoLoader(params))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if !result.Valid() {
			return echo.NewHTTPError(http.StatusBadRequest, middleware.JsonSchemaFormat(result.Errors()))
		}
		return next(c)
	}
}
