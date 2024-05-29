package payload

type CreateMatakuliahRequest struct {
	Name  string `json:"name" form:"name"`
	Dosen string `json:"dosen" form:"dosen"`
}

type CreateMatakuliahResponse struct {
	MatakuliahID uint `json:"matakuliah_id"`
}

type UpdateMatakuliahRequest struct {
	Name  string `json:"name" form:"name"`
	Dosen string `json:"dosen" form:"dosen"`
}

type UpdateMatakuliahResponse struct {
	MatakuliahID uint `json:"matakuliah_id"`
}

type GetMatakuliahResponse struct {
	MatakuliahID uint   `json:"matakuliah_id"`
	Name         string `json:"name" form:"name"`
	Dosen        string `json:"dosen" form:"dosen"`
}

type GetMatakuliahsResponse struct {
	Matakuliahs []GetMatakuliahResponse `json:"matakuliahs"`
}
