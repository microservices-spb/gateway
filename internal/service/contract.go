package service

type DbRepo interface {
	Write(x, y int64)
}
