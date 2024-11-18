package repository

type Repository struct {
	db map[int64]int64
}

func New() *Repository {
	db := make(map[int64]int64)
	return &Repository{db: db}
}

func (r *Repository) Write(x, y int64) {
	r.db[x] = y
}
