package server

import (
	"kawanishi/first_boiler/api_srv/infrastracture/persistence"
)

type Server struct {
	repo *persistence.Repositories
}

func NewServer(r *persistence.Repositories) (*Server, error) {
	return &Server{
		repo: r,
	}, nil
}
