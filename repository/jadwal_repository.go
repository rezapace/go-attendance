package repository

import (
	"context"

	"presensee_project/model"
	"presensee_project/model/payload"
)

type JadwalRepository interface {
	CreateJadwal(ctx context.Context, absen *model.Jadwal) error
	GetSingleJadwal(ctx context.Context, absenID uint) (*model.Jadwal, error)
	GetPageJadwals(ctx context.Context, limit int, offset int) (*model.Jadwals, error)
	GetFilterJadwals(ctx context.Context, limit int, offset int, filter *payload.JadwalFilter) (*model.Jadwals, error)
	UpdateJadwal(ctx context.Context, absen *model.Jadwal) error
	DeleteJadwal(ctx context.Context, absenID uint) error
	CountJadwalFilter(ctx context.Context, filter *payload.JadwalFilter) (int64, error)
	CountJadwal(ctx context.Context) (int64, error)
}
