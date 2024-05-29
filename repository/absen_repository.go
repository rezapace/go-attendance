package repository

import (
	"context"

	"presensee_project/model"
	"presensee_project/model/payload"
)

type AbsenRepository interface {
	CreateAbsen(ctx context.Context, absen *model.Absen) error
	GetSingleAbsen(ctx context.Context, absenID uint) (*model.Absen, error)
	GetPageAbsens(ctx context.Context, limit int, offset int) (*model.Absens, error)
	GetFilterAbsens(ctx context.Context, limit int, offset int, filter *payload.AbsenFilter) (*model.Absens, error)
	UpdateAbsen(ctx context.Context, absen *model.Absen) error
	DeleteAbsen(ctx context.Context, absenID uint) error
	CountAbsenFilter(ctx context.Context, filter *payload.AbsenFilter) (int64, error)
	CountRiwayatMatakuliah(ctx context.Context, filter *payload.AbsenFilter) (int64, error)
	CountAbsen(ctx context.Context) (int64, error)
	GetMatakuliah(ctx context.Context) (*model.Matakuliahs, error)
}
