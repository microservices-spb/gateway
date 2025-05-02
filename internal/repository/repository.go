package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	//"github.com/microservices-spb/gateway/internal/api"
	"github.com/microservices-spb/gateway/internal/model"
)

type Repository struct {
	db map[int64]int64
}

type PostgresUserRepository struct {
	Conn *sqlx.DB
}

func ConnectToDB() *PostgresUserRepository {
	fmt.Println("Connecting to DB")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5433, "master", "master", "usersInfoDB")

	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	return &PostgresUserRepository{Conn: conn}
}

func (r *PostgresUserRepository) SaveUser(ctx context.Context, user *model.RequestData) (string, error) {
	query := "INSERT INTO usersInfo (username, password) VALUES ($1, $2) RETURNING id"
	var id int64
	err := r.Conn.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&id)
	if err != nil {
		return query, err
	}

	return query, nil
}

func (r *PostgresUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	query := "SELECT id, username, password FROM userinfo id=$1"
	err := r.Conn.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) CheckUserInDB(ctx context.Context, username model.User) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM userinfo WHERE username = $1)"
	var exists bool
	err := r.Conn.QueryRowContext(ctx, query, username.Username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, err
}

//var _ api.UserRepository = (*PostgresUserRepository)(nil)

func New() *Repository {
	db := make(map[int64]int64)
	return &Repository{db: db}
}

func (r *Repository) Write(x, y int64) {
	r.db[x] = y
}
