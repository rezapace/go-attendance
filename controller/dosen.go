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
func GetDosensController(c echo.Context) error {
	dosens, err := usecase.GetListDosens()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	totalCount := len(dosens)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"dosens":      dosens,
		"total_count": totalCount,
	})
}

// GetDosenController returns dosen data based on ID
func GetDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dosen, err := usecase.GetDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosen":  dosen,
	})
}

// CreateDosenController creates a new dosen
func CreateDosenController(c echo.Context) error {
	requestPayload := new(payload.CreateDosenRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the user_id exists in the users table
	// _, err := usecase.GetSingleUser(requestPayload.UserID)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	// }

	dosen := &model.Dosen{
		Name:   requestPayload.Name,
		Email:  requestPayload.Email,
		NIP:    requestPayload.NIP,
		Phone:  requestPayload.Phone,
		Image:  requestPayload.Image,
		UserID: requestPayload.UserID,
	}

	err := usecase.CreateDosen(dosen)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateDosenResponse{
		DosenID: dosen.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"dosen":  responsePayload,
	})
}

// UpdateDosenController updates dosen data based on ID
func UpdateDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	dosenToUpdate, err := usecase.GetDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedDosen := new(payload.UpdateDosenRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedDosen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedDosen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update dosen data
	dosenToUpdate.Name = updatedDosen.Name
	dosenToUpdate.Email = updatedDosen.Email
	dosenToUpdate.NIP = updatedDosen.NIP
	dosenToUpdate.Phone = updatedDosen.Phone
	dosenToUpdate.Image = updatedDosen.Image
	dosenToUpdate.UserID = updatedDosen.UserID

	err = usecase.UpdateDosen(&dosenToUpdate) // Pass the pointer to dosenToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateDosenResponse{
		DosenID: dosenToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosen":  response,
	})
}

// DeleteDosenController deletes dosen data based on ID
func DeleteDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Dosen deleted successfully",
	})
}
