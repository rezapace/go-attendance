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

type JadwalRepositoryImpl struct {
	db *gorm.DB
}

func NewJadwalRepositoryImpl(db *gorm.DB) repository.JadwalRepository {
	jadwalRepository := &JadwalRepositoryImpl{
		db: db,
	}

	return jadwalRepository
}

func (u *JadwalRepositoryImpl) CreateJadwal(ctx context.Context, jadwal *model.Jadwal) error {
	err := u.db.WithContext(ctx).Create(jadwal).Error
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

func (u *JadwalRepositoryImpl) GetSingleJadwal(ctx context.Context, jadwalID uint) (*model.Jadwal, error) {
	var jadwal model.Jadwal

	err := u.db.WithContext(ctx).
		Where("id = ?", jadwalID).
		First(&jadwal).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrItemNotFound
		}

		return nil, err
	}

	return &jadwal, nil
}

func (u *JadwalRepositoryImpl) GetPageJadwals(ctx context.Context, limit int, offset int) (*model.Jadwals, error) {
	var jadwals model.Jadwals
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&jadwals).Error
	if err != nil {
		return nil, err
	}

	if len(jadwals) == 0 {
		return nil, utils.ErrItemNotFound
	}

	return &jadwals, nil
}

func (u *JadwalRepositoryImpl) GetFilterJadwals(ctx context.Context, limit int, offset int, filter *payload.JadwalFilter) (*model.Jadwals, error) {
	var jadwals model.Jadwals
	db := u.db.WithContext(ctx)

	// Mengatur filter berdasarkan parameter yang ada
	if filter.ID != 0 {
		db = db.Where("id = ?", filter.ID)
	}

	if filter.Name != "" {
		db = db.Where("name = ?", filter.Name)
	}

	if !filter.JamMulaiAfter.IsZero() {
		db = db.Where("jam_mulai > ?", filter.JamMulaiAfter)
	}

	if !filter.JamMulaiBefore.IsZero() {
		db = db.Where("jam_mulai < ?", filter.JamMulaiBefore)
	}

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	if filter.DosenID != 0 {
		db = db.Where("dosen_id = ?", filter.DosenID)
	}

	if filter.MatakuliahID != 0 {
		db = db.Where("matakuliah_id = ?", filter.MatakuliahID)
	}

	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&jadwals).Error

	if err != nil {
		return nil, err
	}

	if len(jadwals) == 0 {
		return nil, utils.ErrItemNotFound
	}

	return &jadwals, nil
}

func (u *JadwalRepositoryImpl) UpdateJadwal(ctx context.Context, jadwal *model.Jadwal) error {
	result := u.db.WithContext(ctx).Model(&model.Jadwal{}).Where("id = ?", jadwal.ID).Updates(jadwal)
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
		return utils.ErrItemNotFound
	}

	return nil
}

func (d *JadwalRepositoryImpl) DeleteJadwal(ctx context.Context, jadwalID uint) error {
	result := d.db.WithContext(ctx).
		Select("Jadwal").
		Delete(&model.Jadwal{}, "id = ?", jadwalID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrItemNotFound
	}

	return nil
}

func (u *JadwalRepositoryImpl) CountJadwal(ctx context.Context) (int64, error) {
	db, err := u.db.DB()
	if err != nil {
		return 0, err
	}

	query := sq.Select("COUNT(*)").
		From("jadwals a").
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

func (u *JadwalRepositoryImpl) CountJadwalFilter(ctx context.Context, filter *payload.JadwalFilter) (int64, error) {
	db, err := u.db.DB()
	if err != nil {
		return 0, err
	}

	query := sq.Select("COUNT(*)").
		From("jadwals a").
		Where("a.deleted_at IS NULL")

	// Mengatur filter berdasarkan parameter yang ada
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}

	if filter.Name != "" {
		query = query.Where("name = ?", filter.Name)
	}

	if !filter.JamMulaiAfter.IsZero() {
		query = query.Where("jam > ?", filter.JamMulaiAfter)
	}

	if !filter.JamMulaiBefore.IsZero() {
		query = query.Where("jam < ?", filter.JamMulaiBefore)
	}

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.DosenID != 0 {
		query = query.Where("dosen_id = ?", filter.DosenID)
	}

	if filter.MatakuliahID != 0 {
		query = query.Where("matakuliah_id = ?", filter.MatakuliahID)
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
