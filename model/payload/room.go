package payload

type CreateRoomRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	Code     string `json:"code" form:"code" validate:"required"`
}

type CreateRoomResponse struct {
	RoomID uint `json:"room_id"`
}

type UpdateRoomRequest struct {
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Code     string `json:"code" form:"code"`
}

type UpdateRoomResponse struct {
	RoomID uint `json:"room_id"`
}

type GetRoomResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Seat     string `json:"code"`
}

type GetRoomsResponse struct {
	Rooms []GetRoomResponse `json:"rooms"`
}
