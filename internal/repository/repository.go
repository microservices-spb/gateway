package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db   map[int64]int64
	Conn *sqlx.DB
}

func ConnectToDB() *Repository {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "master", "master", "usersInfoDB")

	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	return &Repository{Conn: conn}
}

func New() *Repository {
	db := make(map[int64]int64)
	return &Repository{db: db}
}

func (r *Repository) Write(x, y int64) {
	r.db[x] = y
}
