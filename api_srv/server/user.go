package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"go.uber.org/zap"
)

func (s *Server) User(ctx context.Context, userId int) error {
	user, err := s.repo.User.GetUser(userId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil
		default:
			return errors.WithStack(err)
		}
	}

	logger, err := zap.NewProduction()
	logger.Debug(fmt.Sprintf("%v", user))

	return nil
}
