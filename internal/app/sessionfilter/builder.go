package sessionfilter

import "github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"

type Builder interface {
	BuildCriteria(filter *SessionFilter) *session.Criteria
}

type builder struct{}

func NewBuilder() Builder {
	return builder{}
}

func (b builder) BuildCriteria(filter *SessionFilter) *session.Criteria {
	return &session.Criteria{
		UserID:           filter.UserID,
		OnlyActive:       filter.OnlyActive,
		DeviceType:       filter.DeviceType,
		DeviceName:       filter.DeviceName,
		AppType:          filter.AppType,
		AppVersion:       filter.AppVersion,
		OS:               filter.OS,
		OSVersion:        filter.OSVersion,
		DeviceModel:      filter.DeviceModel,
		IPAddr:           filter.IPAddr,
		LastActiveBefore: filter.LastActiveBefore,
		LastActiveAfter:  filter.LastActiveAfter,
		UpdatedBefore:    filter.UpdatedBefore,
		UpdatedAfter:     filter.UpdatedAfter,
		CreatedBefore:    filter.CreatedBefore,
		CreatedAfter:     filter.CreatedAfter,
	}
}
