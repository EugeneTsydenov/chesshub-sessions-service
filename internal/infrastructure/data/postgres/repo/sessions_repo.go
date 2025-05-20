package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres"
	postgreserrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/pkg/spec"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type PostgresSessionRepositoryImpl struct {
	database *postgres.Database
}

func NewPostgresSessionRepository(db *postgres.Database) *PostgresSessionRepositoryImpl {
	return &PostgresSessionRepositoryImpl{
		database: db,
	}
}

func (r *PostgresSessionRepositoryImpl) Create(ctx context.Context, entity *entity.Session) (*entity.Session, error) {
	query := `INSERT INTO sessions (user_id, ip_address, device_info, expired_at) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id, user_id, ip_address, device_info, is_active, expired_at, updated_at, created_at`

	row := r.database.Pool().QueryRow(ctx, query, entity.UserID(), entity.IPAddr(), entity.DeviceInfo(), entity.ExpiredAt())

	createdSession, err := scanSession(ctx, row)

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, errors.New("insert session to postgres took too long")
	}

	if errors.Is(err, context.Canceled) {
		return nil, errors.New("insert session was canceled by client or system")
	}

	if err != nil {
		return nil, postgreserrors.NewUnresolvedError("failed to create session", err)
	}

	return createdSession, nil
}

func scanSession(_ context.Context, row pgx.Row) (*entity.Session, error) {
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

	b := entity.NewSessionBuilder()
	s := b.
		WithID(id).
		WithUserID(userId).
		WithIPAddr(ipAddr).
		WithDeviceInfo(deviceInfo).
		WithIsActive(isActive).
		WithExpiredAt(expiredAt).
		WithUpdatedAt(updatedAt).
		WithCreatedAt(createdAt).
		Build()

	return s, nil
}

func (r *PostgresSessionRepositoryImpl) GetByID(ctx context.Context, id string) (*entity.Session, error) {
	query := `SELECT * FROM sessions WHERE id = $1 LIMIT 1`
	row := r.database.Pool().QueryRow(ctx, query, id)

	session, err := scanSession(ctx, row)
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, errors.New("get session by id from postgres took too long")
	}

	if errors.Is(err, context.Canceled) {
		return nil, errors.New("get session by id was canceled by client or system")
	}

	if err != nil {
		return nil, postgreserrors.NewUnresolvedError("failed to get session by id", err)
	}

	return session, nil
}

func (r *PostgresSessionRepositoryImpl) GetAll(ctx context.Context, spec spec.Spec) ([]*entity.Session, error) {
	query, args, err := spec.ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := r.database.Pool().Query(ctx, query, args...)

	if err != nil {
		return nil, postgreserrors.NewUnresolvedError("failed to getting sessions", err)
	}
	defer rows.Close()

	var sessions []*entity.Session

	for rows.Next() {
		s, err := scanSession(ctx, rows)
		if err != nil {
			return nil, postgreserrors.NewUnresolvedError("failed to scan session", err)
		}
		sessions = append(sessions, s)
	}

	if err = rows.Err(); err != nil {
		return nil, postgreserrors.NewUnresolvedError("rows iteration error", err)
	}

	return sessions, nil
}

func (r *PostgresSessionRepositoryImpl) Update(ctx context.Context, spec spec.Spec) (*entity.Session, error) {
	query, args, err := spec.ToSQL()
	log.Print(query, args)
	if err != nil {
		return nil, err
	}

	row := r.database.Pool().QueryRow(ctx, query, args...)

	session, err := scanSession(ctx, row)
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, errors.New("update session to postgres took too long")
	}

	if errors.Is(err, context.Canceled) {
		return nil, errors.New("update session was canceled by client or system")
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, postgreserrors.NewNoRowsError("session not found")
	}

	if err != nil {
		return nil, postgreserrors.NewUnresolvedError("failed to update session", err)
	}

	return session, nil
}
