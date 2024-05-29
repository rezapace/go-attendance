package bootsrapper

import (
	"time"

	ControllerPkg "presensee_project/controller"
	RepositoryPkg "presensee_project/repository/impl"
	routes "presensee_project/route"
	ServicePkg "presensee_project/usecase/impl"
	jwtPkg "presensee_project/utils/jwt_service/impl"
	passwordPkg "presensee_project/utils/password/impl"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB, conf map[string]string) {
	passwordFunc := passwordPkg.NewPasswordFuncImpl()
	jwtService := jwtPkg.NewJWTService(conf["JWT_SECRET"], 1*time.Hour)

	// User
	userRepository := RepositoryPkg.NewUserRepositoryImpl(db)
	userService := ServicePkg.NewUserServiceImpl(userRepository, passwordFunc, jwtService)
	userController := ControllerPkg.NewUserController(userService, jwtService)

	// Jadwal
	jadwalRepository := RepositoryPkg.NewJadwalRepositoryImpl(db)
	jadwalService := ServicePkg.NewJadwalServiceImpl(jadwalRepository)
	jadwalController := ControllerPkg.NewJadwalController(jadwalService, jwtService)

	// Absen
	absenRepository := RepositoryPkg.NewAbsenRepositoryImpl(db)
	absenService := ServicePkg.NewAbsenServiceImpl(absenRepository)
	absenController := ControllerPkg.NewAbsenController(absenService, jwtService)

	route := routes.NewRoutes(userController, absenController, jadwalController)
	route.Init(e, conf)
}
