package usecase

import (
	"context"

	"presensee_project/model/payload"
)

type JadwalService interface {
	CreateJadwal(ctx context.Context, jadwal *payload.CreateJadwalRequest) error
	GetSingleJadwal(ctx context.Context, jadwalID uint) (*payload.GetSingleJadwalResponse, error)
	GetPageJadwals(ctx context.Context, page int, limit int) (*payload.GetPageJadwalsResponse, int64, error)
	GetFilterJadwals(ctx context.Context, page int, limit int, filter *payload.JadwalFilter) (*payload.GetPageJadwalsResponse, int64, error)
	UpdateJadwal(ctx context.Context, jadwalID uint, request *payload.UpdateJadwalRequest) error
	DeleteJadwal(ctx context.Context, jadwalID uint) error
}
