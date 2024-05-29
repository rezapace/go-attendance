package config

import (
	"fmt"
	"log"
	"os"
	"presensee_project/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadConfig() map[string]string {
	env := make(map[string]string)

	env["DB_HOST"] = os.Getenv("DB_HOST")
	env["DB_PORT"] = os.Getenv("DB_PORT")
	env["DB_USER"] = os.Getenv("DB_USER")
	env["DB_PASS"] = os.Getenv("DB_PASS")
	env["DB_NAME"] = os.Getenv("DB_NAME")
	env["JWT_SECRET"] = os.Getenv("JWT_SECRET")
	return env
}

func InitDB() *gorm.DB {
	env := LoadConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		env["DB_USER"],
		env["DB_PASS"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	local, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		log.Fatal("failed to load timezone: ", err)
	}
	time.Local = local

	InitMigrate()
	return DB
}

func InitMigrate() {

	DB.AutoMigrate(&model.User{}, &model.Mahasiswa{}, &model.Dosen{}, &model.Jadwal{}, &model.Matakuliah{}, &model.Room{}, &model.Absen{}, &model.Jurusan{})
}
