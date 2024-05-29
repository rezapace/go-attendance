package controller

import (
	"net/http"

	"presensee_project/model/payload"
	"presensee_project/usecase"
	"presensee_project/utils"
	"presensee_project/utils/jwt_service"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type JadwalController struct {
	jadwalService usecase.JadwalService
	jwtService    jwt_service.JWTService
}

func NewJadwalController(jadwalService usecase.JadwalService, jwtService jwt_service.JWTService) *JadwalController {
	return &JadwalController{
		jadwalService: jadwalService,
		jwtService:    jwtService,
	}
}

func (u *JadwalController) CreateJadwal(c echo.Context) error {
	jadwal := new(payload.CreateJadwalRequest)
	if err := c.Bind(jadwal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(jadwal); err != nil {
		return err
	}

	err := u.jadwalService.CreateJadwal(c.Request().Context(), jadwal)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating jadwal",
	})
}

func (u *JadwalController) UpdateJadwal(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	jadwal := new(payload.UpdateJadwalRequest)
	if err := c.Bind(jadwal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(jadwal); err != nil {
		return err
	}

	err := u.jadwalService.UpdateJadwal(c.Request().Context(), jadwal.ID, jadwal)
	if err != nil {
		switch err {
		case utils.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update jadwal",
	})
}

func (u *JadwalController) GetSingleJadwal(c echo.Context) error {
	jadwalIDParam := c.Param("jadwal_id")
	jadwalID64, err := strconv.ParseUint(jadwalIDParam, 10, 0)
	jadwalID := uint(jadwalID64)

	jadwal, err := u.jadwalService.GetSingleJadwal(c.Request().Context(), jadwalID)
	if err != nil {
		if err == utils.ErrItemNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)

	switch {
	case role == "pegawai":
		fallthrough
	case role == "admin" || role == "Mahasiswa" || role == "Dosen":
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success getting user",
			"data":    jadwal,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *JadwalController) GetPageJadwal(c echo.Context) error {

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	jadwal, count, err := u.jadwalService.GetPageJadwals(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrItemNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    jadwal,
		"meta": echo.Map{
			"total": count,
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (u *JadwalController) GetFilterJadwal(c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	userIDStr := c.QueryParam("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	}

	dosenIDStr := c.QueryParam("dosen_id")
	dosenID, err := strconv.ParseUint(dosenIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid dosen_id")
	}

	jadwalIDStr := c.QueryParam("jadwal_id")
	jadwalID, err := strconv.ParseUint(jadwalIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid jadwal_id")
	}

	matakuliahIDStr := c.QueryParam("matakuliah_id")
	matakuliahID, err := strconv.ParseUint(matakuliahIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid matakuliah_id")
	}

	roomIDStr := c.QueryParam("room_id")
	roomID, err := strconv.ParseUint(roomIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid room_id")
	}

	jamAfterStr := c.QueryParam("jam_after")
	jamAfterTime, err := time.Parse(time.RFC3339, jamAfterStr)

	jamBeforeStr := c.QueryParam("jam_before")
	jamBeforeTime, err := time.Parse(time.RFC3339, jamBeforeStr)

	filter := payload.JadwalFilter{
		UserID:         uint(userID),
		DosenID:        uint(dosenID),
		MatakuliahID:   uint(matakuliahID),
		RoomID:         uint(roomID),
		Name:           c.QueryParam("name"),
		ID:             uint(jadwalID),
		JamMulaiAfter:  jamAfterTime,
		JamMulaiBefore: jamBeforeTime,
	}

	jadwal, count, err := u.jadwalService.GetFilterJadwals(c.Request().Context(), int(pageInt), int(limitInt), &filter)
	if err != nil {
		if err == utils.ErrItemNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting jadwal",
		"data":    jadwal,
		"meta": echo.Map{
			"total": count,
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *JadwalController) DeleteJadwal(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	jadwalIDParam := c.Param("jadwal_id")
	jadwalID64, err := strconv.ParseUint(jadwalIDParam, 10, 0)
	jadwalID := uint(jadwalID64)
	err = d.jadwalService.DeleteJadwal(c.Request().Context(), jadwalID)
	if err != nil {
		switch err {
		case utils.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting jadwal",
	})
}
