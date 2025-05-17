package port

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/spec"
)

type SessionsRepo interface {
	Create(ctx context.Context, entity *entity.Session) (*entity.Session, error)
	GetById(ctx context.Context, id string) (*entity.Session, error)
	GetAll(ctx context.Context, spec spec.Spec) ([]*entity.Session, error)
}
