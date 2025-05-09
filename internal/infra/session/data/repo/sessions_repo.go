package repo

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data/dberrors"
	"github.com/jackc/pgx/v5"
	"log"
	"time"

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

func (r *SessionsRepoImpl) Create(ctx context.Context, model *session.Session) (*session.Session, error) {
	query := `INSERT INTO sessions (user_id, ip_address, device_info, expired_at) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id, user_id, ip_address, device_info, is_active, expired_at, updated_at, created_at`

	row := r.pool.QueryRow(ctx, query, model.UserId(), model.IpAddr(), model.DeviceInfo(), model.ExpiredAt())

	createdSession, err := scanSession(ctx, row)

	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return nil, err
	}

	if err != nil {
		log.Println("Error creating session:", err)
		return nil, dberrors.NewUnresolvedError("failed to create session")
	}

	return createdSession, nil
}

func scanSession(ctx context.Context, row pgx.Row) (*session.Session, error) {
	var (
		id         string
		userId     int64
		ipAddr     string
		deviceInfo string
		isActive   bool
		expiredAt  time.Time
		updatedAt  time.Time
		createdAt  time.Time
	)

	err := row.Scan(&id, &userId, &ipAddr, &deviceInfo, &isActive, &expiredAt, &updatedAt, &createdAt)
	if err != nil {
		return nil, err
	}

	b := session.NewBuilder()
	s := b.
		WithId(id).
		WithUserId(userId).
		WithIpAddr(ipAddr).
		WithDeviceInfo(deviceInfo).
		WithIsActive(isActive).
		WithExpiredAt(expiredAt).
		WithUpdatedAt(updatedAt).
		WithCreatedAt(createdAt).
		Build()

	return s, nil
}

func (r *SessionsRepoImpl) Read(_ context.Context, spec any) (*session.Session, error) {
	return nil, nil
}

func (r *SessionsRepoImpl) Update(_ context.Context, model *session.Session) (*session.Session, error) {
	return model, nil
}

func (r *SessionsRepoImpl) Delete(_ context.Context, model *session.Session) (*session.Session, error) {
	return model, nil
}
