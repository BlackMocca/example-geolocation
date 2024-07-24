package repository

import (
	"context"

	"github.com/BlackMocca/sqlx"
	"github.com/Blackmocca/example-geolocation/models"
	"github.com/Blackmocca/example-geolocation/service/tracking"
)

type psqlRepository struct {
	client *sqlx.DB
}

func NewPsqlRepository(client *sqlx.DB) tracking.Repository {
	return &psqlRepository{client: client}
}

func (p psqlRepository) SaveTracking(ctx context.Context, tracking *models.Tracking) error {
	var tx, err = p.client.Beginx()
	if err != nil {
		panic(err)
	}

	sql := `INSERT INTO trackings (id, plate_number, lat, lon, track_time)
	VALUES (nextval('trackings_id_seq'), ?, ?, ?, ?)
	`
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)

	stmt, err := tx.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, tracking.PlateNumber, tracking.Lat, tracking.Lon, tracking.Tracktime); err != nil {
		return err
	}

	return tx.Commit()
}
