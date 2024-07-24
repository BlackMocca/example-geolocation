package models

import (
	"github.com/Blackmocca/utils"
)

type Tracking struct {
	Id          int              `json:"id"`
	PlateNumber string           `json:"plate_number"`
	Lat         float64          `json:"lat"`
	Lon         float64          `json:"lon"`
	Tracktime   *utils.Timestamp `json:"track_time"`
}
