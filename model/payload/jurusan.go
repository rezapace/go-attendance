package payload

type CreateJurusanRequest struct {
	Name     string `json:"name" form:"name"`
	Fakultas string `json:"fakultas" form:"fakultas"`
}

type CreateJurusanResponse struct {
	JurusanID uint `json:"jurusan_id"`
}

type UpdateJurusanRequest struct {
	Name     string `json:"name" form:"name"`
	Fakultas string `json:"fakultas" form:"fakultas"`
}

type UpdateJurusanResponse struct {
	JurusanID uint `json:"jurusan"`
}

type GetJurusanResponse struct {
	JurusanID uint   `json:"jurusan_id"`
	Name      string `json:"name" form:"name"`
	Fakultas  string `json:"fakultas" form:"fakultas"`
}

type GetJurusansResponse struct {
	Jurusans []GetJurusanResponse `json:"jurusans"`
}
