package payload

type CreateDosenRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required,email"`
	NIP    string `json:"nip" form:"nip" validate:"required"`
	Phone  string `json:"phone" form:"phone"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
}

type CreateDosenResponse struct {
	DosenID uint `json:"dosen_id"`
}

type UpdateDosenRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required,email"`
	NIP    string `json:"nip" form:"nip" validate:"required"`
	Phone  string `json:"phone" form:"phone"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
}

type UpdateDosenResponse struct {
	DosenID uint `json:"dosen_id"`
}

type GetDosenResponse struct {
	DosenID uint   `json:"dosen_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	NIP     string `json:"nip"`
	Image   string `json:"image"`
	UserID  uint   `json:"user_id"`
}

type GetDosensResponse struct {
	Dosens []GetDosenResponse `json:"dosens"`
}
