package postgres

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/Masterminds/squirrel"
	"time"
)

type (
	SessionQueryFactory interface {
		BuildQuery(criteria *session.Criteria) (string, []interface{}, error)
	}
	sessionQueryFactory struct {
		psql squirrel.StatementBuilderType
	}
)

func NewSessionQueryFactory() SessionQueryFactory {
	return &sessionQueryFactory{
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (f *sessionQueryFactory) BuildQuery(criteria *session.Criteria) (string, []interface{}, error) {
	query := f.psql.
		Select(
			"id",
			"user_id",
			"device_type",
			"device_name",
			"app_type",
			"app_version",
			"os",
			"os_version",
			"device_model",
			"ip_address",
			"city",
			"country",
			"is_active",
			"lifetime",
			"last_active_at",
			"updated_at",
			"created_at",
		).
		From("sessions")

	query = f.applyFilters(query, criteria)

	sql, args, err := query.ToSql()
	return sql, args, err
}

func (f *sessionQueryFactory) applyFilters(query squirrel.SelectBuilder, criteria *session.Criteria) squirrel.SelectBuilder {
	if criteria == nil {
		return query
	}

	if criteria.UserID != nil {
		query = query.Where(squirrel.Eq{"user_id": *criteria.UserID})
	}

	if criteria.OnlyActive != nil && *criteria.OnlyActive {
		query = query.Where(squirrel.Eq{"is_active": true})
		query = query.Where("last_active_at + lifetime > ?", time.Now())
	}

	if criteria.DeviceType != nil {
		query = query.Where(squirrel.Eq{"device_type": *criteria.DeviceType})
	}

	if criteria.DeviceName != nil {
		query = query.Where(squirrel.Eq{"device_name": *criteria.DeviceName})
	}

	if criteria.AppType != nil {
		query = query.Where(squirrel.Eq{"app_type": *criteria.AppType})
	}

	if criteria.AppVersion != nil {
		query = query.Where(squirrel.Eq{"app_version": *criteria.AppVersion})
	}

	if criteria.OS != nil {
		query = query.Where(squirrel.Eq{"os": *criteria.OS})
	}

	if criteria.OSVersion != nil {
		query = query.Where(squirrel.Eq{"os_version": *criteria.OSVersion})
	}

	if criteria.DeviceModel != nil {
		query = query.Where(squirrel.Eq{"device_model": *criteria.DeviceModel})
	}

	if criteria.IPAddr != nil {
		query = query.Where(squirrel.Eq{"ip_address": *criteria.IPAddr})
	}

	if criteria.LastActiveBefore != nil {
		query = query.Where(squirrel.LtOrEq{"last_active_at": *criteria.LastActiveBefore})
	}

	if criteria.LastActiveAfter != nil {
		query = query.Where(squirrel.GtOrEq{"last_active_at": *criteria.LastActiveAfter})
	}

	if criteria.UpdatedBefore != nil {
		query = query.Where(squirrel.LtOrEq{"updated_at": *criteria.UpdatedBefore})
	}

	if criteria.UpdatedAfter != nil {
		query = query.Where(squirrel.GtOrEq{"updated_at": *criteria.UpdatedAfter})
	}

	if criteria.CreatedBefore != nil {
		query = query.Where(squirrel.LtOrEq{"created_at": *criteria.CreatedBefore})
	}

	if criteria.CreatedAfter != nil {
		query = query.Where(squirrel.GtOrEq{"created_at": *criteria.CreatedAfter})
	}

	return query
}
