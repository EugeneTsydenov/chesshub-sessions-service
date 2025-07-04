package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/redis"
	"github.com/google/uuid"
	r "github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

var (
	DefaultSessionTTL = session.DefaultSessionLifetime
	SessionPrefix     = "session"
)

var (
	ErrSessionNotFound   = errors.New("session not found in cache")
	ErrInvalidSession    = errors.New("invalid session data")
	ErrEmptyRedisHash    = errors.New("redis hash is empty")
	ErrMissingField      = errors.New("missing required field")
	ErrInvalidFieldValue = errors.New("invalid field value")
)

const (
	fieldID           = "id"
	fieldUserID       = "user_id"
	fieldIsActive     = "is_active"
	fieldLifetime     = "lifetime"
	fieldLastActiveAt = "last_active_at"
	fieldCreatedAt    = "created_at"
	fieldUpdatedAt    = "updated_at"
	fieldDeviceType   = "device_type"
	fieldAppType      = "app_type"
	fieldAppVersion   = "app_version"
	fieldOS           = "os"
	fieldIPAddr       = "ip_addr"
	fieldDeviceName   = "device_name"
	fieldOSVersion    = "os_version"
	fieldDeviceModel  = "device_model"
	fieldCity         = "city"
	fieldCountry      = "country"
)

var requiredFields = []string{
	fieldID, fieldUserID, fieldIsActive, fieldLifetime,
	fieldLastActiveAt, fieldCreatedAt, fieldUpdatedAt,
	fieldDeviceType, fieldAppType, fieldAppVersion,
	fieldOS, fieldIPAddr,
}

type SessionCache struct {
	database *redis.Database
	prefix   string
	ttl      time.Duration
}

func NewSessionCache(database *redis.Database) *SessionCache {
	return &SessionCache{
		database: database,
		prefix:   SessionPrefix,
		ttl:      DefaultSessionTTL,
	}
}

func (c *SessionCache) HSet(ctx context.Context, s *session.Session) error {
	if s == nil {
		return ErrInvalidSession
	}

	hash, err := c.buildSessionHash(s)
	if err != nil {
		return fmt.Errorf("failed to build session hash: %w", err)
	}

	key := c.buildKey(s.ID())
	log.Printf("Setting cache key: %s", key) // ДОБАВИТЬ
	log.Printf("Hash data: %+v", hash)       // ДОБАВИТЬ

	pipe := c.database.Client().Pipeline()
	pipe.HSet(ctx, key, hash)

	ttl := s.Lifetime()
	if ttl <= 0 {
		ttl = c.ttl
	}
	pipe.Expire(ctx, key, ttl)

	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("Failed to exec pipeline: %v", err) // ДОБАВИТЬ
		return fmt.Errorf("failed to save session to database: %w", err)
	}

	log.Printf("Successfully saved session with key: %s", key) // ДОБАВИТЬ
	return nil
}

func (c *SessionCache) HGet(ctx context.Context, id uuid.UUID) (*session.Session, error) {
	key := c.buildKey(id)

	result, err := c.database.Client().HGetAll(ctx, key).Result()
	if err != nil {
		if errors.Is(err, r.Nil) {
			return nil, ErrSessionNotFound
		}
		return nil, fmt.Errorf("failed to get session from database: %w", err)
	}

	if len(result) == 0 {
		return nil, ErrSessionNotFound
	}

	s, err := c.buildSessionFromHash(result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse session from database: %w", err)
	}

	return s, nil
}

func (c *SessionCache) Del(ctx context.Context, id uuid.UUID) error {
	key := c.buildKey(id)

	result := c.database.Client().Del(ctx, key)
	if err := result.Err(); err != nil {
		return fmt.Errorf("failed to delete session from database: %w", err)
	}

	return nil
}

func (c *SessionCache) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	key := c.buildKey(id)

	result, err := c.database.Client().Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check session existence: %w", err)
	}

	return result > 0, nil
}

func (c *SessionCache) ExtendTTL(ctx context.Context, id uuid.UUID, ttl time.Duration) error {
	key := c.buildKey(id)

	result := c.database.Client().Expire(ctx, key, ttl)
	if err := result.Err(); err != nil {
		return fmt.Errorf("failed to extend session TTL: %w", err)
	}

	if !result.Val() {
		return ErrSessionNotFound
	}

	return nil
}

func (c *SessionCache) buildKey(id uuid.UUID) string {
	return fmt.Sprintf("%s:%s", c.prefix, id.String())
}

func (c *SessionCache) buildSessionHash(s *session.Session) (map[string]interface{}, error) {
	deviceInfo := s.DeviceInfo()
	if deviceInfo == nil {
		return nil, fmt.Errorf("%w: device info is nil", ErrInvalidSession)
	}

	location := s.Location()
	if location == nil {
		return nil, fmt.Errorf("%w: location is nil", ErrInvalidSession)
	}

	hash := map[string]interface{}{
		fieldID:           s.ID().String(),
		fieldUserID:       s.UserID(),
		fieldIsActive:     s.IsActive(),
		fieldLifetime:     s.Lifetime().Nanoseconds(),
		fieldLastActiveAt: s.LastActiveAt().Unix(),
		fieldCreatedAt:    s.CreatedAt().Unix(),
		fieldUpdatedAt:    s.UpdatedAt().Unix(),
		fieldDeviceType:   int(deviceInfo.DeviceType()),
		fieldAppType:      int(deviceInfo.AppType()),
		fieldAppVersion:   deviceInfo.AppVersion(),
		fieldOS:           deviceInfo.OS(),
		fieldIPAddr:       deviceInfo.IPAddr(),
	}

	if deviceInfo.DeviceName() != nil {
		hash[fieldDeviceName] = *deviceInfo.DeviceName()
	}
	if deviceInfo.OSVersion() != nil {
		hash[fieldOSVersion] = *deviceInfo.OSVersion()
	}
	if deviceInfo.DeviceModel() != nil {
		hash[fieldDeviceModel] = *deviceInfo.DeviceModel()
	}

	if location.City() != nil {
		hash[fieldCity] = *location.City()
	}
	if location.Country() != nil {
		hash[fieldCountry] = *location.Country()
	}

	return hash, nil
}

func (c *SessionCache) buildSessionFromHash(hash map[string]string) (*session.Session, error) {
	if len(hash) == 0 {
		return nil, ErrEmptyRedisHash
	}

	if err := c.validateRequiredFields(hash); err != nil {
		return nil, err
	}

	builder := session.NewBuilder()

	if err := c.parseMainFields(builder, hash); err != nil {
		return nil, err
	}

	if err := c.parseTimeFields(builder, hash); err != nil {
		return nil, err
	}

	deviceInfo, err := c.parseDeviceInfo(hash)
	if err != nil {
		return nil, err
	}
	builder.WithDeviceInfo(deviceInfo)

	location := c.parseLocation(hash)
	builder.WithLocation(location)

	return builder.Build(), nil
}

func (c *SessionCache) validateRequiredFields(hash map[string]string) error {
	for _, field := range requiredFields {
		if _, exists := hash[field]; !exists {
			return fmt.Errorf("%w: %s", ErrMissingField, field)
		}
	}
	return nil
}

func (c *SessionCache) parseMainFields(builder *session.Builder, hash map[string]string) error {
	id, err := uuid.Parse(hash[fieldID])
	if err != nil {
		return fmt.Errorf("%w: invalid session id: %v", ErrInvalidFieldValue, err)
	}
	builder.WithID(id)

	userID, err := strconv.ParseInt(hash[fieldUserID], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: invalid user_id: %v", ErrInvalidFieldValue, err)
	}
	builder.WithUserID(userID)

	isActive, err := strconv.ParseBool(hash[fieldIsActive])
	if err != nil {
		return fmt.Errorf("%w: invalid is_active: %v", ErrInvalidFieldValue, err)
	}
	builder.WithIsActive(isActive)

	lifetimeNanos, err := strconv.ParseInt(hash[fieldLifetime], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: invalid lifetime: %v", ErrInvalidFieldValue, err)
	}
	builder.WithLifetime(time.Duration(lifetimeNanos))

	return nil
}

func (c *SessionCache) parseTimeFields(builder *session.Builder, hash map[string]string) error {
	lastActiveUnix, err := strconv.ParseInt(hash[fieldLastActiveAt], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: invalid last_active_at: %v", ErrInvalidFieldValue, err)
	}
	builder.WithLastActiveAt(time.Unix(lastActiveUnix, 0))

	createdUnix, err := strconv.ParseInt(hash[fieldCreatedAt], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: invalid created_at: %v", ErrInvalidFieldValue, err)
	}
	builder.WithCreatedAt(time.Unix(createdUnix, 0))

	updatedUnix, err := strconv.ParseInt(hash[fieldUpdatedAt], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: invalid updated_at: %v", ErrInvalidFieldValue, err)
	}
	builder.WithUpdatedAt(time.Unix(updatedUnix, 0))

	return nil
}

func (c *SessionCache) parseDeviceInfo(hash map[string]string) (*session.DeviceInfo, error) {
	deviceType, err := strconv.Atoi(hash[fieldDeviceType])
	if err != nil {
		return nil, fmt.Errorf("%w: invalid device_type: %v", ErrInvalidFieldValue, err)
	}

	appType, err := strconv.Atoi(hash[fieldAppType])
	if err != nil {
		return nil, fmt.Errorf("%w: invalid app_type: %v", ErrInvalidFieldValue, err)
	}

	var deviceName, osVersion, deviceModel *string

	if value, exists := hash[fieldDeviceName]; exists && value != "" {
		deviceName = &value
	}
	if value, exists := hash[fieldOSVersion]; exists && value != "" {
		osVersion = &value
	}
	if value, exists := hash[fieldDeviceModel]; exists && value != "" {
		deviceModel = &value
	}

	return session.NewDeviceInfo(
		session.DeviceType(deviceType),
		deviceName,
		session.AppType(appType),
		hash[fieldAppVersion],
		hash[fieldOS],
		osVersion,
		deviceModel,
		hash[fieldIPAddr],
	), nil
}

func (c *SessionCache) parseLocation(hash map[string]string) *session.Location {
	var city, country *string

	if value, exists := hash[fieldCity]; exists && value != "" {
		city = &value
	}
	if value, exists := hash[fieldCountry]; exists && value != "" {
		country = &value
	}

	return session.NewLocation(city, country)
}
