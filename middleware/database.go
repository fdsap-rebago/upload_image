package middleware

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	Err    error
)

func PostgreSQLConnect() {

	DSN := ("host=34.87.184.130 user=postgres password=p0stgR35_t3sT_p@55_2022 dbname=mfs port=5432 sslmode=disable timezone=Asia/Manila")
	DBConn, Err = gorm.Open(postgres.Open(DSN), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	DBConn.AutoMigrate()
}
