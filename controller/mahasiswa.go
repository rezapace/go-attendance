package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload" // Ubah import path ini
	usecase "presensee_project/usecase/impl"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMahasiswasController returns all mahasiswa data
func GetMahasiswasController(c echo.Context) error {
	mahasiswas, err := usecase.GetListMahasiswas()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	totalCount := len(mahasiswas)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"mahasiswas":  mahasiswas,
		"total_count": totalCount,
	})
}

// GetMahasiswaController returns mahasiswa data based on ID
func GetMahasiswaController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	mahasiswa, err := usecase.GetMahasiswa(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"mahasiswa": mahasiswa,
	})
}

// CreateMahasiswaController creates a new mahasiswa
func CreateMahasiswaController(c echo.Context) error {
	requestPayload := new(payload.CreateMahasiswaRequest)

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

	mahasiswa := &model.Mahasiswa{
		Name:       requestPayload.Name,
		Email:      requestPayload.Email,
		NIM:        requestPayload.NIM,
		Image:      requestPayload.Image,
		Phone:      requestPayload.Phone,
		Jurusan:    requestPayload.Jurusan,
		Fakultas:   requestPayload.Fakultas,
		TahunMasuk: requestPayload.TahunMasuk,
		IPK:        requestPayload.IPK,
		UserID:     requestPayload.UserID,
	}

	err := usecase.CreateMahasiswa(mahasiswa)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateMahasiswaResponse{
		MahasiswaID: mahasiswa.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":    "success",
		"mahasiswa": responsePayload,
	})
}

// UpdateMahasiswaController updates mahasiswa data based on ID
func UpdateMahasiswaController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	mahasiswaToUpdate, err := usecase.GetMahasiswa(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedMahasiswa := new(payload.UpdateMahasiswaRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedMahasiswa); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedMahasiswa); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update mahasiswa data
	mahasiswaToUpdate.Name = updatedMahasiswa.Name
	mahasiswaToUpdate.Email = updatedMahasiswa.Email
	mahasiswaToUpdate.NIM = updatedMahasiswa.NIM
	mahasiswaToUpdate.Image = updatedMahasiswa.Image
	mahasiswaToUpdate.Phone = updatedMahasiswa.Phone
	mahasiswaToUpdate.Jurusan = updatedMahasiswa.Jurusan
	mahasiswaToUpdate.Fakultas = updatedMahasiswa.Fakultas
	mahasiswaToUpdate.TahunMasuk = updatedMahasiswa.TahunMasuk
	mahasiswaToUpdate.IPK = updatedMahasiswa.IPK
	mahasiswaToUpdate.UserID = updatedMahasiswa.UserID

	err = usecase.UpdateMahasiswa(&mahasiswaToUpdate) // Pass the pointer to mahasiswaToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateMahasiswaResponse{
		MahasiswaID: mahasiswaToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"mahasiswa": response,
	})
}

// DeleteMahasiswaController deletes mahasiswa data based on ID
func DeleteMahasiswaController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteMahasiswa(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Mahasiswa deleted successfully",
	})
}
