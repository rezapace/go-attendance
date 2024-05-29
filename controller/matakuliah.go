package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload" // Ubah import path ini
	usecase "presensee_project/usecase/impl"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GetMatakuliahsController returns all matakuliah data
func GetMatakuliahsController(c echo.Context) error {
	matakuliahs, err := usecase.GetListMatakuliahs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"matakuliahs": matakuliahs,
	})
}

// GetMatakuliahByNameAndDateController returns matakuliah data based on name and date
func GetMatakuliahByNameAndDateController(c echo.Context) error {
	name := c.QueryParam("name")
	dateParam := c.QueryParam("date")

	date, err := time.Parse(time.RFC3339, dateParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format")
	}

	matakuliah, err := usecase.GetMatakuliahByNameAndDate(name, date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := payload.GetMatakuliahResponse{
		MatakuliahID: matakuliah.ID,
		Name:         matakuliah.Name,
		Dosen:        matakuliah.Dosen,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"matakuliah": response,
	})
}

// GetMatakuliahController returns matakuliah data based on ID
func GetMatakuliahController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	matakuliah, err := usecase.GetMatakuliah(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"matakuliah": matakuliah,
	})
}

// CreateMatakuliahController creates a new matakuliah
func CreateMatakuliahController(c echo.Context) error {
	requestPayload := new(payload.CreateMatakuliahRequest)

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

	matakuliah := &model.Matakuliah{
		Name:  requestPayload.Name,
		Dosen: requestPayload.Dosen,
	}

	err := usecase.CreateMatakuliah(matakuliah)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateMatakuliahResponse{
		MatakuliahID: matakuliah.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":     "success",
		"matakuliah": responsePayload,
	})
}

// UpdateMatakuliahController updates matakuliah data based on ID
func UpdateMatakuliahController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	matakuliahToUpdate, err := usecase.GetMatakuliah(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedMatakuliah := new(payload.UpdateMatakuliahRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedMatakuliah); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedMatakuliah); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update matakuliah data
	matakuliahToUpdate.Name = updatedMatakuliah.Name
	matakuliahToUpdate.Dosen = updatedMatakuliah.Dosen

	err = usecase.UpdateMatakuliah(&matakuliahToUpdate) // Pass the pointer to matakuliahToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateMatakuliahResponse{
		MatakuliahID: matakuliahToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"matakuliah": response,
	})
}

// DeleteMatakuliahController deletes matakuliah data based on ID
func DeleteMatakuliahController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteMatakuliah(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Matakuliah deleted successfully",
	})
}
