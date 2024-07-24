package tracking

import (
	"context"

	"github.com/Blackmocca/example-geolocation/models"
)

type Usecase interface {
	SaveTracking(ctx context.Context, tracking *models.Tracking) error
}
