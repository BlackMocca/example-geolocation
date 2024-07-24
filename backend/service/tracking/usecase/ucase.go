package usecase

import (
	"context"

	"github.com/Blackmocca/example-geolocation/models"
	"github.com/Blackmocca/example-geolocation/service/tracking"
)

type usecase struct {
	trackingRepo tracking.Repository
}

func NewUsecase(trackingRepo tracking.Repository) tracking.Usecase {
	return &usecase{trackingRepo: trackingRepo}
}

func (u usecase) SaveTracking(ctx context.Context, tracking *models.Tracking) error {
	if err := u.trackingRepo.SaveTracking(ctx, tracking); err != nil {
		return err
	}

	// processing lastest location here

	return nil
}
