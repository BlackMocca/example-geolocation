package tracking

import (
	"context"

	"github.com/Blackmocca/example-geolocation/models"
)

type Repository interface {
	SaveTracking(ctx context.Context, tracking *models.Tracking) error
}
