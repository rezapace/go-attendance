package routes

import (
	"presensee_project/controller"
	"presensee_project/utils/validation"

	ControllerPkg "presensee_project/controller"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	userController   *ControllerPkg.UserController
	absenController  *ControllerPkg.AbsenController
	jadwalController *ControllerPkg.JadwalController
}

func NewRoutes(userController *ControllerPkg.UserController, absenController *ControllerPkg.AbsenController, jadwalController *ControllerPkg.JadwalController) *Routes {
	return &Routes{
		userController:   userController,
		absenController:  absenController,
		jadwalController: jadwalController,
	}
}

func (r *Routes) Init(e *echo.Echo, conf map[string]string) {
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf["JWT_SECRET"]),
	})

	v1 := e.Group("/v1")

	// Users
	users := v1.Group("/users")
	users.POST("/signup/", r.userController.SignUpUser)
	users.POST("/login/", r.userController.LoginUser)
	users.POST("/admin/", r.userController.LoginAdmin)
	users.POST("/dosen/", r.userController.LoginDosen)

	usersWithAuth := users.Group("", jwtMiddleware)
	usersWithAuth.GET("/", r.userController.GetBriefUsers)
	usersWithAuth.PUT("/", r.userController.UpdateUser)
	usersWithAuth.GET("/:user_id/", r.userController.GetSingleUser, jwtMiddleware)
	usersWithAuth.DELETE("/:user_id/", r.userController.DeleteUser, jwtMiddleware)

	// mahasiswa collection
	mahasiswa := v1.Group("/mahasiswa", jwtMiddleware)
	mahasiswa.GET("/:id/", controller.GetMahasiswaController)
	mahasiswa.POST("/", controller.CreateMahasiswaController)
	mahasiswa.PUT("/:id/", controller.UpdateMahasiswaController)
	mahasiswa.DELETE("/:id/", controller.DeleteMahasiswaController)

	mahasiswaNoJwt := v1.Group("/mahasiswa")
	mahasiswaNoJwt.GET("/", controller.GetMahasiswasController)

	// dosen
	dosen := v1.Group("/dosen", jwtMiddleware)
	dosen.GET("/:id/", controller.GetDosenController)
	dosen.POST("/", controller.CreateDosenController)
	dosen.PUT("/:id/", controller.UpdateDosenController)
	dosen.DELETE("/:id/", controller.DeleteDosenController)

	dosenNoJwt := v1.Group("/dosen")
	dosenNoJwt.GET("/", controller.GetDosensController)

	// room
	room := v1.Group("/room", jwtMiddleware)
	room.GET("/:id/", controller.GetRoomController)
	room.POST("/", controller.CreateRoomController)
	room.PUT("/:id/", controller.UpdateRoomController)
	room.DELETE("/:id/", controller.DeleteRoomController)

	roomNoJwt := v1.Group("/room")
	roomNoJwt.GET("/", controller.GetRoomsController)

	// matakuliah
	matakuliah := v1.Group("/matakuliah", jwtMiddleware)
	matakuliah.GET("/", controller.GetMatakuliahsController)
	matakuliah.GET("/:id/", controller.GetMatakuliahController)
	matakuliah.GET("/:name/", controller.GetMatakuliahByNameAndDateController)
	matakuliah.POST("/", controller.CreateMatakuliahController)
	matakuliah.PUT("/:id/", controller.UpdateMatakuliahController)
	matakuliah.DELETE("/:id/", controller.DeleteMatakuliahController)

	// absen
	absens := v1.Group("/absens", jwtMiddleware)
	absens.POST("/", r.absenController.CreateAbsen)
	absens.PUT("/", r.absenController.UpdateAbsen, jwtMiddleware)
	absens.GET("/:absen_id/", r.absenController.GetSingleAbsen, jwtMiddleware)
	absens.GET("/", r.absenController.GetPageAbsen)
	absens.GET("/riwayat/", r.absenController.GetRiwayat)
	absens.GET("/dashboard/", r.absenController.GetRiwayatDashboard)
	absens.GET("/filter/", r.absenController.GetFilterAbsen)
	absens.DELETE("/:absen_id/", r.absenController.DeleteAbsen, jwtMiddleware)

	// Jurusan
	jurusan := v1.Group("/jurusan", jwtMiddleware)
	jurusan.GET("/", controller.GetJurusansController)
	jurusan.GET("/:id/", controller.GetJurusanController)
	jurusan.POST("/", controller.CreateJurusanController)
	jurusan.PUT("/:id/", controller.UpdateJurusanController)
	jurusan.DELETE("/:id/", controller.DeleteJurusanController)

	// Jadwal
	jadwals := v1.Group("/jadwals", jwtMiddleware)
	jadwals.POST("/", r.jadwalController.CreateJadwal)
	jadwals.PUT("/", r.jadwalController.UpdateJadwal, jwtMiddleware)
	jadwals.GET("/:jadwal_id/", r.jadwalController.GetSingleJadwal, jwtMiddleware)
	jadwals.GET("/", r.jadwalController.GetPageJadwal)
	jadwals.GET("/filter/", r.jadwalController.GetFilterJadwal)
	jadwals.DELETE("/:jadwal_id/", r.jadwalController.DeleteJadwal, jwtMiddleware)

	//Upload File
	upload := v1.Group("/upload")
	upload.POST("/", controller.UploadFile)
}
