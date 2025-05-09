package repo

import (
	"context"
	"log"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data"
)

type SessionsRepo interface {
	repo[*session.Session]
}

type SessionsRepoImpl struct {
	pool data.DbPool
}

var _ SessionsRepo = new(SessionsRepoImpl)

func NewSessionsRepo(pool data.DbPool) *SessionsRepoImpl {
	return &SessionsRepoImpl{pool: pool}
}

func (s *SessionsRepoImpl) Create(_ context.Context, model *session.Session) (*session.Session, error) {
	log.Print("CreateRepo Create called")

	return model, nil
}

func (s *SessionsRepoImpl) Read(_ context.Context, spec any) (*session.Session, error) {
	return nil, nil
}

func (s *SessionsRepoImpl) Update(_ context.Context, model *session.Session) (*session.Session, error) {
	return model, nil
}

func (s *SessionsRepoImpl) Delete(_ context.Context, model *session.Session) (*session.Session, error) {
	return model, nil
}
