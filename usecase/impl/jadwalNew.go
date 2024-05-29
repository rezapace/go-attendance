package usecase

import (
	"context"
	"log"

	"presensee_project/model/payload"
	"presensee_project/repository"
	database "presensee_project/repository/impl"
	"presensee_project/usecase"

	"github.com/google/uuid"
)

type (
	JadwalServiceImpl struct {
		jadwalRepository repository.JadwalRepository
	}
)

func NewJadwalServiceImpl(jadwalRepository repository.JadwalRepository) usecase.JadwalService {
	return &JadwalServiceImpl{
		jadwalRepository: jadwalRepository,
	}
}

func (u *JadwalServiceImpl) CreateJadwal(ctx context.Context, jadwal *payload.CreateJadwalRequest) error {

	jadwalEntity := jadwal.ToEntity()
	jadwalEntity.ID = uint(uuid.New().ID())

	err := u.jadwalRepository.CreateJadwal(ctx, jadwalEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *JadwalServiceImpl) GetSingleJadwal(ctx context.Context, jadwalID uint) (*payload.GetSingleJadwalResponse, error) {
	jadwal, err := d.jadwalRepository.GetSingleJadwal(ctx, jadwalID)
	if err != nil {
		return nil, err
	}

	dosen, err := database.GetDosenByID(jadwal.DosenID)
	if err != nil {
		return nil, err
	}

	matakuliah, err := database.GetMatakuliahByID(jadwal.MatakuliahID)
	if err != nil {
		return nil, err
	}

	room, err := database.GetRoomByID(jadwal.RoomID)
	if err != nil {
		return nil, err
	}

	var jadwalResponse = payload.NewGetSingleJadwalResponse(jadwal)
	jadwalResponse.Matakuliah = matakuliah
	jadwalResponse.Dosen = dosen
	jadwalResponse.Room = room

	return jadwalResponse, nil
}

func (u *JadwalServiceImpl) GetPageJadwals(ctx context.Context, page int, limit int) (*payload.GetPageJadwalsResponse, int64, error) {
	count, err := u.jadwalRepository.CountJadwal(ctx)
	if err != nil {
		log.Println("error while counting jadwals: ", err)
		return nil, 0, err
	}

	if count == 0 {
		return nil, 0, nil
	}

	offset := (page - 1) * limit

	jadwals, err := u.jadwalRepository.GetPageJadwals(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var jadwalResponse = payload.NewGetPageJadwalsResponse(jadwals)

	for i := range *jadwalResponse {
		dosen, err := database.GetDosenByID((*jadwalResponse)[i].DosenID)
		_ = err
		room, err := database.GetRoomByID((*jadwalResponse)[i].RoomID)
		_ = err
		(*jadwalResponse)[i].Dosen = dosen
		(*jadwalResponse)[i].Room = room
	}
	return jadwalResponse, count, nil
}

func (u *JadwalServiceImpl) GetFilterJadwals(ctx context.Context, page int, limit int, filter *payload.JadwalFilter) (*payload.GetPageJadwalsResponse, int64, error) {
	count, err := u.jadwalRepository.CountJadwalFilter(ctx, filter)
	if err != nil {
		log.Println("error while counting jadwals: ", err)
		return nil, 0, err
	}

	if count == 0 {
		return nil, 0, nil
	}

	offset := (page - 1) * limit

	jadwals, err := u.jadwalRepository.GetFilterJadwals(ctx, limit, offset, filter)
	if err != nil {
		return nil, 0, err
	}

	var jadwalResponse = payload.NewGetPageJadwalsResponse(jadwals)

	for i := range *jadwalResponse {
		dosen, err := database.GetDosenByID((*jadwalResponse)[i].DosenID)
		_ = err
		room, err := database.GetRoomByID((*jadwalResponse)[i].RoomID)
		_ = err
		(*jadwalResponse)[i].Dosen = dosen
		(*jadwalResponse)[i].Room = room
	}
	return jadwalResponse, count, nil
}

func (u *JadwalServiceImpl) UpdateJadwal(ctx context.Context, jadwalID uint, updateJadwal *payload.UpdateJadwalRequest) error {
	jadwal := updateJadwal.ToEntity()
	jadwal.ID = jadwalID

	return u.jadwalRepository.UpdateJadwal(ctx, jadwal)
}

func (d *JadwalServiceImpl) DeleteJadwal(ctx context.Context, jadwalID uint) error {

	return d.jadwalRepository.DeleteJadwal(ctx, jadwalID)
}
