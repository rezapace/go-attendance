package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload" // Ubah import path ini

	usecase "presensee_project/usecase/impl"

	"strconv"

	"github.com/labstack/echo/v4"
)

// GetJurusansController returns all jurusan data
func GetJurusansController(c echo.Context) error {
	jurusans, err := usecase.GetListJurusans()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"jurusans": jurusans,
	})
}

// GetJurusanController returns jurusan data based on ID
func GetJurusanController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	jurusan, err := usecase.GetJurusan(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"jurusan": jurusan,
	})
}

// CreateJurusanController creates a new jurusan
func CreateJurusanController(c echo.Context) error {
	requestPayload := new(payload.CreateJurusanRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user_id exists in the users table
	// _, err := usecase.GetUser(requestPayload.UserID)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	// }

	jurusan := &model.Jurusan{
		Name:     requestPayload.Name,
		Fakultas: requestPayload.Fakultas,
	}

	err := usecase.CreateJurusan(jurusan)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateJurusanResponse{
		JurusanID: jurusan.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"jurusan": responsePayload,
	})
}

// UpdateJurusanController updates jurusan data based on ID
func UpdateJurusanController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	jurusanToUpdate, err := usecase.GetJurusan(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedJurusan := new(payload.UpdateJurusanRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedJurusan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedJurusan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update jurusan data
	jurusanToUpdate.Name = updatedJurusan.Name
	jurusanToUpdate.Fakultas = updatedJurusan.Fakultas

	err = usecase.UpdateJurusan(&jurusanToUpdate) // Pass the pointer to jurusanToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateJurusanResponse{
		JurusanID: jurusanToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"jurusan": response,
	})
}

// DeleteJurusanController deletes jurusan data based on ID
func DeleteJurusanController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteJurusan(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Jurusan deleted successfully",
	})
}
