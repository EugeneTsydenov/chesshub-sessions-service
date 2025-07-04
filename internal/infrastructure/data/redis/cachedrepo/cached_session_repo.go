package cachedrepo

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/google/uuid"
	"log"
	"time"
)

type CachedSessionRepo struct {
	sessionCache interfaces.SessionCache
	sessionRepo  interfaces.SessionRepo
}

func NewCachedSessionRepo(cache interfaces.SessionCache, repo interfaces.SessionRepo) interfaces.SessionRepo {
	return &CachedSessionRepo{
		sessionCache: cache,
		sessionRepo:  repo,
	}
}

func (r *CachedSessionRepo) Create(ctx context.Context, s *session.Session) (*session.Session, error) {
	log.Print("[DEBUG] Create session")
	s, err := r.sessionRepo.Create(ctx, s)
	if err != nil {
		return nil, err
	}

	r.saveSessionToCacheAsync(s)

	return s, nil
}

func (r *CachedSessionRepo) GetByID(ctx context.Context, sessionID uuid.UUID) (*session.Session, error) {
	s, err := r.sessionCache.HGet(ctx, sessionID)
	if err != nil {
		log.Printf("Cache miss for session: %s, error: %v", sessionID, err)
		s, err = r.sessionRepo.GetByID(ctx, sessionID)
		if err != nil {
			return nil, err
		}
		log.Printf("Retrieved session from DB: %s", s.ID())

		r.saveSessionToCacheAsync(s)
		return s, nil
	}
	log.Printf("Cache hit for session: %s", sessionID)
	return s, nil
}

func (r *CachedSessionRepo) Find(ctx context.Context, criteria *session.Criteria) ([]*session.Session, error) {
	return r.sessionRepo.Find(ctx, criteria)
}

func (r *CachedSessionRepo) Update(ctx context.Context, s *session.Session) (*session.Session, error) {
	updated, err := r.sessionRepo.Update(ctx, s)
	if err != nil {
		_ = r.sessionCache.Del(ctx, s.ID())

		return nil, err
	}

	r.saveSessionToCacheAsync(updated)

	return updated, nil
}

func (r *CachedSessionRepo) saveSessionToCacheAsync(s *session.Session) {
	go func() {
		bgCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := r.sessionCache.HSet(bgCtx, s); err != nil {
			log.Printf("Failed to cache session: %v", err)
			return
		}

		log.Print("Save session to cache successfully")
	}()
}
