package payload

import "presensee_project/model"

type UserSignUpRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

func (u *UserSignUpRequest) ToEntity() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Role:     u.Role,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Telp     string `json:"telp"`
	Role     string `json:"role"`
}

func (u *UserUpdateRequest) ToEntity() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Role:     u.Role,
	}
}

type GetSingleUserResponse struct {
	ID        uint            `json:"id"`
	Email     string          `json:"email"`
	Name      string          `json:"name"`
	Role      string          `json:"role"`
	Mahasiswa model.Mahasiswa `json:"mahasiswa"`
	Dosen     model.Dosen     `json:"dosen"`
}

func NewGetSingleUserResponse(user *model.User) *GetSingleUserResponse {
	return &GetSingleUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}

type BriefUserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func NewBriefUserResponse(user *model.User) *BriefUserResponse {
	return &BriefUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}

type BriefUsersResponse []BriefUserResponse

func NewBriefUsersResponse(users *model.Users) *BriefUsersResponse {
	var briefUsersResponse BriefUsersResponse
	for _, user := range *users {
		briefUsersResponse = append(briefUsersResponse, *NewBriefUserResponse(&user))
	}
	return &briefUsersResponse
}
