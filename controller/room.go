package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload" // Ubah import path ini
	usecase "presensee_project/usecase/impl"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetDosensController returns all dosen data
func GetRoomsController(c echo.Context) error {
	rooms, err := usecase.GetListRooms()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"rooms":  rooms,
	})
}

// GetDosenController returns dosen data based on ID
func GetRoomController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	room, err := usecase.GetRoom(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"room":   room,
	})
}

// CreateDosenController creates a new dosen
func CreateRoomController(c echo.Context) error {
	requestPayload := new(payload.CreateRoomRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room := &model.Room{
		Name:     requestPayload.Name,
		Location: requestPayload.Location,
		Code:     requestPayload.Code,
	}

	err := usecase.CreateRoom(room)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateRoomResponse{
		RoomID: room.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"room":   responsePayload,
	})
}

// UpdateRoomController updates dosen data based on ID
func UpdateRoomController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	roomToUpdate, err := usecase.GetRoom(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedRoom := new(payload.UpdateRoomRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedRoom); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedRoom); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update dosen data
	roomToUpdate.Name = updatedRoom.Name
	roomToUpdate.Location = updatedRoom.Location
	roomToUpdate.Code = updatedRoom.Code

	err = usecase.UpdateRoom(&roomToUpdate) // Pass the pointer to dosenToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateRoomResponse{
		RoomID: roomToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"room":   response,
	})
}

// DeleteDosenController deletes dosen data based on ID
func DeleteRoomController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteRoom(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Room deleted successfully",
	})
}
