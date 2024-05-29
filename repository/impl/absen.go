package database

import (
	"context"
	"strings"

	"presensee_project/model"
	"presensee_project/model/payload"
	"presensee_project/repository"
	"presensee_project/utils"

	sq "github.com/Masterminds/squirrel"

	"gorm.io/gorm"
)

type AbsenRepositoryImpl struct {
	db *gorm.DB
}

func NewAbsenRepositoryImpl(db *gorm.DB) repository.AbsenRepository {
	absenRepository := &AbsenRepositoryImpl{
		db: db,
	}

	return absenRepository
}

func (u *AbsenRepositoryImpl) CreateAbsen(ctx context.Context, absen *model.Absen) error {
	err := u.db.WithContext(ctx).Create(absen).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *AbsenRepositoryImpl) GetSingleAbsen(ctx context.Context, absenID uint) (*model.Absen, error) {
	var absen model.Absen

	err := u.db.WithContext(ctx).
		Where("id = ?", absenID).
		First(&absen).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrAbsenNotFound
		}

		return nil, err
	}

	return &absen, nil
}

func (u *AbsenRepositoryImpl) GetPageAbsens(ctx context.Context, limit int, offset int) (*model.Absens, error) {
	var absens model.Absens
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&absens).Error
	if err != nil {
		return nil, err
	}

	if len(absens) == 0 {
		return nil, utils.ErrAbsenNotFound
	}

	return &absens, nil
}

func (u *AbsenRepositoryImpl) GetFilterAbsens(ctx context.Context, limit int, offset int, filter *payload.AbsenFilter) (*model.Absens, error) {
	var absens model.Absens
	db := u.db.WithContext(ctx)

	// Mengatur filter berdasarkan parameter yang ada
	if filter.ID != 0 {
		db = db.Where("id = ?", filter.ID)
	}

	if !filter.CreatedAfter.IsZero() {
		db = db.Where("time_attemp > ?", filter.CreatedAfter)
	}

	if !filter.CreatedBefore.IsZero() {
		db = db.Where("time_attemp < ?", filter.CreatedBefore)
	}

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	if filter.MahasiswaID != 0 {
		db = db.Where("mahasiswa_id = ?", filter.MahasiswaID)
	}

	if filter.JadwalID != 0 {
		db = db.Where("jadwal_id = ?", filter.JadwalID)
	}

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	if filter.Matakuliah != "" {
		db = db.Where("matakuliah = ?", filter.Matakuliah)
	}

	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&absens).Error

	if err != nil {
		return nil, err
	}

	if len(absens) == 0 {
		return nil, utils.ErrAbsenNotFound
	}

	return &absens, nil
}

func (u *AbsenRepositoryImpl) UpdateAbsen(ctx context.Context, absen *model.Absen) error {
	result := u.db.WithContext(ctx).Model(&model.Absen{}).Where("id = ?", absen.ID).Updates(absen)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrAbsenNotFound
	}

	return nil
}

func (d *AbsenRepositoryImpl) DeleteAbsen(ctx context.Context, absenID uint) error {
	result := d.db.WithContext(ctx).
		Select("Absen").
		Delete(&model.Absen{}, "id = ?", absenID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrAbsenNotFound
	}

	return nil
}

func (u *AbsenRepositoryImpl) CountAbsen(ctx context.Context) (int64, error) {
	db, err := u.db.DB()
	if err != nil {
		return 0, err
	}

	query := sq.Select("COUNT(*)").
		From("absens a").
		Where("a.deleted_at IS NULL")

	rows, err := query.RunWith(db).QueryContext(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close()
	}()

	var count int64
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (u *AbsenRepositoryImpl) CountAbsenFilter(ctx context.Context, filter *payload.AbsenFilter) (int64, error) {
	db, err := u.db.DB()
	if err != nil {
		return 0, err
	}

	query := sq.Select("COUNT(*)").
		From("absens a").
		Where("a.deleted_at IS NULL")

	// Mengatur filter berdasarkan parameter yang ada
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}

	if !filter.CreatedAfter.IsZero() {
		query = query.Where("created_at > ?", filter.CreatedAfter)
	}

	if !filter.CreatedBefore.IsZero() {
		query = query.Where("created_at < ?", filter.CreatedBefore)
	}

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.MahasiswaID != 0 {
		query = query.Where("mahasiswa_id = ?", filter.MahasiswaID)
	}

	if filter.JadwalID != 0 {
		query = query.Where("jadwal_id = ?", filter.JadwalID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Matakuliah != "" {
		query = query.Where("matakuliah = ?", filter.Matakuliah)
	}

	rows, err := query.RunWith(db).QueryContext(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close()
	}()

	var count int64
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (u *AbsenRepositoryImpl) CountRiwayatMatakuliah(ctx context.Context, filter *payload.AbsenFilter) (int64, error) {
	db, err := u.db.DB()
	if err != nil {
		return 0, err
	}

	print("count")

	query := sq.Select("COUNT(*)").
		From("absens a").
		Where("a.deleted_at IS NULL")

	// Mengatur filter berdasarkan parameter yang ada
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}

	if !filter.CreatedAfter.IsZero() {
		query = query.Where("created_at > ?", filter.CreatedAfter)
	}

	if !filter.CreatedBefore.IsZero() {
		query = query.Where("created_at < ?", filter.CreatedBefore)
	}

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.MahasiswaID != 0 {
		query = query.Where("mahasiswa_id = ?", filter.MahasiswaID)
	}

	if filter.JadwalID != 0 {
		query = query.Where("jadwal_id = ?", filter.JadwalID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Matakuliah != "" {
		query = query.Where("matakuliah = ?", filter.Matakuliah)
	}

	rows, err := query.RunWith(db).QueryContext(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close()
	}()

	var count int64
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (u *AbsenRepositoryImpl) GetMatakuliah(ctx context.Context) (*model.Matakuliahs, error) {
	var matakuliahs model.Matakuliahs
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&matakuliahs).Error
	if err != nil {
		return nil, err
	}

	if len(matakuliahs) == 0 {
		return nil, utils.ErrAbsenNotFound
	}

	return &matakuliahs, nil
}
