package pkg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbRespository struct {
	Db *gorm.DB
}

var repo *DbRespository

func GetDatabaseConn() *DbRespository {
	if repo == nil {
		var err error

		db, err := gorm.Open(postgres.Open(os.Getenv("DB_DNS")))
		if err != nil {
			panic(err)
		}

		stat, err := db.DB()
		if err != nil {
			panic(err)
		}

		db_max_idle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
		db_max_con, _ := strconv.Atoi(os.Getenv("DB_MAX_CON"))

		stat.SetMaxIdleConns(db_max_idle)
		stat.SetMaxOpenConns(db_max_con)
		stat.SetConnMaxLifetime(time.Hour)

		repo = &DbRespository{
			Db: db,
		}

		fmt.Println("Postgres connection successfully")
	}

	return repo
}

func (r *DbRespository) PingDatabase() {
	stat, err := r.Db.DB()
	if err != nil {
		panic(err)
	}

	if err := stat.Ping(); err != nil {
		panic(err)
	}
}