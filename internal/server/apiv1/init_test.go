package apiv1

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var h *Handler

func TestMain(m *testing.M) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3307)/rawati_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	h = &Handler{
		Environment: "testing",
		DB:          db,
	}
	m.Run()
}
