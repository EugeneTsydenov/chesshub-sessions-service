package repo

import (
	"context"
	"errors"
	"fmt"
	domainerrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres"
	postgreserrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres/errors"
	"time"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PostgresSessionRepo struct {
	database *postgres.Database

	queryFactory postgres.SessionQueryFactory
}

var _ interfaces.SessionRepo = new(PostgresSessionRepo)

func NewPostgresSessionRepository(db *postgres.Database, factory postgres.SessionQueryFactory) *PostgresSessionRepo {
	return &PostgresSessionRepo{database: db, queryFactory: factory}
}

func (r *PostgresSessionRepo) Create(ctx context.Context, s *session.Session) (*session.Session, error) {
	query := `INSERT INTO sessions (
				id, user_id, device_type, device_name, app_type, 
                app_version, os, os_version, device_model, ip_address, 
                city, country, is_active, lifetime, last_active_at
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10, 
			    $11, $12, $13, $14, $15
			) RETURNING 
			    id,
    user_id,
    device_type,
    device_name,
    app_type,
    app_version,
    os,
    os_version,
    device_model,
    ip_address,
    city,
    country,
    is_active,
    lifetime,
    last_active_at,
    updated_at,
    created_at`

	deviceInfo := s.DeviceInfo()
	location := s.Location()

	row := r.database.Pool().QueryRow(ctx, query,
		s.ID(),
		s.UserID(),
		deviceInfo.DeviceType(),
		deviceInfo.DeviceName(),
		deviceInfo.AppType(),
		deviceInfo.AppVersion(),
		deviceInfo.OS(),
		deviceInfo.OSVersion(),
		deviceInfo.DeviceModel(),
		deviceInfo.IPAddr(),
		location.City(),
		location.Country(),
		s.IsActive(),
		s.Lifetime(),
		s.LastActiveAt(),
	)

	s, err := scanSession(row)
	if err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Create", err, nil)
	}

	return s, nil
}

func (r *PostgresSessionRepo) GetByID(ctx context.Context, sessionID uuid.UUID) (*session.Session, error) {
	query := `
	SELECT 
		id,
		user_id,
		device_type,
		device_name,
		app_type,
		app_version,
		os,
		os_version,
		device_model,
		ip_address,
		city,
		country,
		is_active,
		lifetime,
		last_active_at,
		updated_at,
		created_at
	FROM sessions
	WHERE id = $1
	`

	row := r.database.Pool().QueryRow(ctx, query, sessionID)
	s, err := scanSession(row)
	if err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.GetByID", err, func(e error) error {
			if errors.Is(e, pgx.ErrNoRows) {
				return domainerrors.ErrSessionNotFound
			}
			return fmt.Errorf("scan error: %w", e)
		})
	}

	return s, nil
}

func (r *PostgresSessionRepo) Update(ctx context.Context, session *session.Session) (*session.Session, error) {
	query := `
	UPDATE sessions
	SET
		user_id = $1,
		device_type = $2,
		device_name = $3,
		app_type = $4,
		app_version = $5,
		os = $6,
		os_version = $7,
		device_model = $8,
		ip_address = $9,
		city = $10,
		country = $11,
		is_active = $12,
		lifetime = $13,
		last_active_at = $14,
		updated_at = $15,
		created_at = $16
	WHERE id = $17
	RETURNING 
		id,
		user_id,
		device_type,
		device_name,
		app_type,
		app_version,
		os,
		os_version,
		device_model,
		ip_address,
		city,
		country,
		is_active,
		lifetime,
		last_active_at,
		updated_at,
		created_at
	`

	deviceInfo := session.DeviceInfo()
	location := session.Location()

	row := r.database.Pool().QueryRow(ctx, query,
		session.UserID(),
		deviceInfo.DeviceType(),
		deviceInfo.DeviceName(),
		deviceInfo.AppType(),
		deviceInfo.AppVersion(),
		deviceInfo.OS(),
		deviceInfo.OSVersion(),
		deviceInfo.DeviceModel(),
		deviceInfo.IPAddr(),
		location.City(),
		location.Country(),
		session.IsActive(),
		session.Lifetime(),
		session.LastActiveAt(),
		session.UpdatedAt(),
		session.CreatedAt(),
		session.ID(),
	)

	s, err := scanSession(row)

	if err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Update", err, nil)
	}

	return s, nil
}

func (r *PostgresSessionRepo) Find(ctx context.Context, criteria *session.Criteria) ([]*session.Session, error) {
	query, args, err := r.queryFactory.BuildQuery(criteria)
	if err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Find query builder", err, nil)
	}

	rows, err := r.database.Pool().Query(ctx, query, args...)
	if err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Find query", err, nil)
	}
	defer rows.Close()

	var sessions []*session.Session

	for rows.Next() {
		s, err := scanSession(rows)
		if err != nil {
			return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Find scan row", err, nil)
		}
		sessions = append(sessions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, postgreserrors.WrapWithMapper("PostgresSessionRepo.Find rows iteration", err, nil)
	}

	return sessions, nil
}

func scanSession(row pgx.Row) (*session.Session, error) {
	var (
		id           uuid.UUID
		userID       int64
		deviceType   session.DeviceType
		deviceName   *string
		appType      session.AppType
		appVersion   string
		os           string
		osVersion    *string
		deviceModel  *string
		ipAddr       string
		city         *string
		country      *string
		isActive     bool
		lifetime     time.Duration
		lastActiveAt time.Time
		updatedAt    time.Time
		createdAt    time.Time
	)

	err := row.Scan(
		&id,
		&userID,
		&deviceType,
		&deviceName,
		&appType,
		&appVersion,
		&os,
		&osVersion,
		&deviceModel,
		&ipAddr,
		&city,
		&country,
		&isActive,
		&lifetime,
		&lastActiveAt,
		&updatedAt,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	deviceInfo := session.NewDeviceInfo(
		deviceType,
		deviceName,
		appType,
		appVersion,
		os,
		osVersion,
		deviceModel,
		ipAddr,
	)

	location := session.NewLocation(city, country)

	s := session.NewBuilder().
		WithID(id).
		WithUserID(userID).
		WithDeviceInfo(deviceInfo).
		WithLocation(location).
		WithIsActive(isActive).
		WithLifetime(lifetime).
		WithLastActiveAt(lastActiveAt).
		WithCreatedAt(createdAt).
		WithUpdatedAt(updatedAt).
		Build()

	return s, nil
}
