package service

type Service struct {
	dbR DbRepo
}

func New(dbR DbRepo) *Service {
	return &Service{
		dbR: dbR,
	}
}

func (s *Service) Mulity(x, y int64) int64 {
	res := x * y
	s.dbR.Write(x, y)
	return res
}
