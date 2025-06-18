package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres"
	postgreserrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PostgresSessionRepo struct {
	database *postgres.Database
}

var _ interfaces.SessionRepo = new(PostgresSessionRepo)

func NewPostgresSessionRepository(db *postgres.Database) *PostgresSessionRepo {
	return &PostgresSessionRepo{database: db}
}

func (r *PostgresSessionRepo) Create(ctx context.Context, s *session.Session) (*uuid.UUID, error) {
	query := `INSERT INTO sessions (
				id, user_id, device_type, device_name, app_type, 
                app_version, os, os_version, device_model, ip_address, 
                city, country, is_active, last_active_at
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10, 
			    $11, $12, $13, $14
			) RETURNING id`

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
		s.LastActiveAt(),
	)

	var id *uuid.UUID

	err := row.Scan(&id)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, fmt.Errorf("PostgresSessionRepo.Create ctx=%v", ctxErr)
		}

		return nil, postgreserrors.NewUnresolvedError("creation session finished with error", err)
	}

	return id, nil
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
		last_active_at,
		updated_at,
		created_at,
		lifetime
	FROM sessions
	WHERE id = $1
	`

	row := r.database.Pool().QueryRow(ctx, query, sessionID)
	s, err := scanSession(row)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, fmt.Errorf("PostgresSessionRepo.GetByID ctx=%v", ctxErr)
		}

		return nil, postgreserrors.NewUnresolvedError("getting session by id finished with error", err)
	}

	return s, nil
}

func (r *PostgresSessionRepo) Update(ctx context.Context, session *session.Session) (*uuid.UUID, error) {
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
	RETURNING id
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

	var id *uuid.UUID

	err := row.Scan(&id)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, fmt.Errorf("PostgresSessionRepo.Update ctx=%v", ctxErr)
		}

		return nil, postgreserrors.NewUnresolvedError("updating session finished with error", err)
	}

	return id, nil
}

//func (r *PostgresSessionRepo) GetByID(ctx context.Context, id string) (*session.Session, error) {
//	query := `SELECT * FROM sessions WHERE id = $1`
//	row := r.database.Pool().QueryRow(ctx, query, id)
//
//	s, err := scanSession(row)
//
//	if ctxErr := ctx.Err(); ctxErr != nil {
//		return nil, fmt.Errorf("PostgresSessionRepo.GetByID: %w", ctxErr)
//	}
//
//	if err != nil {
//		return nil, postgreserrors.NewUnresolvedError("retrieving session finished with error", err)
//	}
//
//	return s, nil
//}
//
//func (r *PostgresSessionRepo) GetActiveSessions(ctx context.Context, userID int64) ([]*session.Session, error) {
//	query := `SELECT * FROM sessions WHERE user_id = $1 AND is_active = true AND expired_at > now()`
//	rows, err := r.database.Pool().Query(ctx, query, userID)
//
//	if ctxErr := ctx.Err(); ctxErr != nil {
//		return nil, fmt.Errorf("PostgresSessionRepo.GetActiveSessions: %w", ctxErr)
//	}
//
//	if err != nil {
//		return nil, postgreserrors.NewUnresolvedError("failed to get active sessions", err)
//	}
//	defer rows.Close()
//
//	var sessions []*session.Session
//	for rows.Next() {
//		s, err := scanSession(rows)
//		if err != nil {
//			return nil, postgreserrors.NewUnresolvedError("failed to scan session on get active sessions", err)
//		}
//		sessions = append(sessions, s)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, postgreserrors.NewUnresolvedError("rows iteration error on get active sessions", err)
//	}
//
//	return sessions, nil
//}
//
//func (r *PostgresSessionRepo) Update(ctx context.Context, s *session.Session) error {
//	query := `UPDATE sessions
//			  SET
//				user_id = $1, device_type = $2, device_name = $3, app_name = $4, app_version = $5,
//			    os = $6, os_version = $7, device_model = $8, ip_address = $9, city = $10,
//			    country = $11, is_active = $12, last_active_at = $13, expired_at = $14,
//			    updated_at = $15
//			  WHERE id = $16`
//
//	_, err := r.database.Pool().Exec(ctx, query,
//		s.GetUserID(),
//		s.GetDeviceType(),
//		s.GetDeviceName(),
//		s.GetAppName(),
//		s.GetAppVersion(),
//		s.GetOS(),
//		s.GetOSVersion(),
//		s.GetDeviceModel(),
//		s.GetIPAddr(),
//		s.GetCity(),
//		s.GetCountry(),
//		s.GetIsActive(),
//		s.GetLastActiveAt(),
//		s.GetExpiredAt(),
//		s.GetUpdatedAt(),
//		s.GetID(),
//	)
//
//	if ctxErr := ctx.Err(); ctxErr != nil {
//		return fmt.Errorf("PostgresSessionRepo.Update: %w", ctxErr)
//	}
//
//	if err != nil {
//		return postgreserrors.NewUnresolvedError("update session finished with error", err)
//	}
//
//	return nil
//}
//
//func (r *PostgresSessionRepo) Delete(ctx context.Context, s *session.Session) error {
//	query := `DELETE FROM sessions WHERE id = $1`
//
//	_, err := r.database.Pool().Exec(ctx, query, s.GetID())
//
//	if ctxErr := ctx.Err(); ctxErr != nil {
//		return fmt.Errorf("PostgresSessionRepo.Delete: %w", ctxErr)
//	}
//
//	if err != nil {
//		return postgreserrors.NewUnresolvedError("failed to delete session", err)
//	}
//
//	return nil
//}

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
		&lastActiveAt,
		&updatedAt,
		&createdAt,
		&lifetime,
	)
	if err != nil {
		return nil, err
	}

	deviceInfo := session.NewDeviceInfo(deviceType, deviceName, appType, appVersion, os, osVersion, deviceModel, ipAddr)
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
