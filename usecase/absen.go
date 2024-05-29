package usecase

import (
	"context"

	"presensee_project/model/payload"
)

type AbsenService interface {
	CreateAbsen(ctx context.Context, absen *payload.CreateAbsenRequest) error
	GetSingleAbsen(ctx context.Context, absenID uint) (*payload.GetSingleAbsenResponse, error)
	GetPageAbsens(ctx context.Context, page int, limit int) (*payload.GetPageAbsensResponse, int64, error)
	GetFilterAbsens(ctx context.Context, page int, limit int, filter *payload.AbsenFilter) (*payload.GetPageAbsensResponse, int64, error)
	CountRiwayatMatakuliah(ctx context.Context, filter *payload.AbsenFilter) (int64, error)
	UpdateAbsen(ctx context.Context, absenID uint, request *payload.UpdateAbsenRequest) error
	DeleteAbsen(ctx context.Context, absenID uint) error
	GetMatakuliah(ctx context.Context) (*payload.GetPageMatakuliahssResponse, error)
}
