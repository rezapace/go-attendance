package payload

type CreateMahasiswaRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email" validate:"required,email"`
	NIM        string `json:"nim" form:"nim" validate:"required"`
	Image      string `json:"image" form:"image"`
	Phone      string `json:"phone" form:"phone"`
	Jurusan    string `json:"jurusan" form:"jurusan"`
	Fakultas   string `json:"fakultas" form:"fakultas"`
	TahunMasuk string `json:"tahun_masuk" form:"tahun_masuk"`
	IPK        string `json:"ipk" form:"ipk"`
	UserID     uint   `json:"user_id" form:"user_id" validate:"required"`
}

type CreateMahasiswaResponse struct {
	MahasiswaID uint `json:"mahasiswa_id"`
}

type UpdateMahasiswaRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	NIM        string `json:"nim" form:"nim"`
	Image      string `json:"image" form:"image"`
	Phone      string `json:"phone" form:"phone"`
	Jurusan    string `json:"jurusan" form:"jurusan"`
	Fakultas   string `json:"fakultas" form:"fakultas"`
	TahunMasuk string `json:"tahun_masuk" form:"tahun_masuk"`
	IPK        string `json:"ipk" form:"ipk"`
	UserID     uint   `json:"user_id" form:"user_id"`
}

type UpdateMahasiswaResponse struct {
	MahasiswaID uint `json:"mahasiswa"`
}

type GetMahasiswaResponse struct {
	MahasiswaID uint   `json:"mahasiswa_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	NIM         string `json:"nim"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
}

type GetMahasiswasResponse struct {
	Mahasiswas []GetMahasiswaResponse `json:"mahasiswas"`
}
