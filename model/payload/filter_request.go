package payload

import "time"

type AbsenFilter struct {
	ID            uint      `query:"absen_id" validate:"omitempty, uint"`
	CreatedAfter  time.Time `query:"created_after"`
	CreatedBefore time.Time `query:"created_before"`
	UserID        uint      `query:"user_id" validate:"omitempty, uint"`
	MahasiswaID   uint      `query:"mahasiswa_id" validate:"omitempty, uint"`
	JadwalID      uint      `query:"jadwal_id" validate:"omitempty, uint"`
	Status        string    `query:"status" validate:"omitempty, string"`
	Matakuliah    string    `query:"matakuliah" validate:"omitempty, string"`
	IsKonfirmasi  bool      `query:"is_konfirmasi"`
}

type JadwalFilter struct {
	ID             uint      `query:"id"`
	MatakuliahID   uint      `query:"matakuliah_id"`
	RoomID         uint      `query:"room_id"`
	Sks            string    `query:"sks"`
	JamMulaiAfter  time.Time `query:"jam_after"`
	JamMulaiBefore time.Time `query:"jam_before"`
	Name           string    `query:"name"`
	Description    string    `query:"description"`
	UserID         uint      `query:"user_id" form:"user_id"`
	DosenID        uint      `query:"dosen_id" form:"dosen_id"`
}
