package sessionspec

import (
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/errors"
	"strings"
	"time"
)

type SessionUpdateSpec struct {
	sessionId string
	filterMap map[string]string
}

func NewSessionUpdateSpec(sessionId string, filterMap map[string]string) *SessionUpdateSpec {
	return &SessionUpdateSpec{
		sessionId: sessionId,
		filterMap: filterMap,
	}
}

func (s *SessionUpdateSpec) ToSQL() (string, []any, error) {
	if len(s.filterMap) < 1 {
		return "", nil, errors.NewInvalidFieldError("field map should not be empty")
	}

	baseQuery := "UPDATE sessions SET"
	var updates []string
	var args []any
	paramIndex := 1

	for key, value := range s.filterMap {
		column, ok := fieldToColumn[key]
		if !ok {
			return "", nil, errors.NewInvalidFieldError(fmt.Sprintf("invalid filter field: '%s' is not recognized", key))
		}

		updates = append(updates, fmt.Sprintf("%s=$%d", column, paramIndex))
		args = append(args, value)
		paramIndex++
	}

	updates = append(updates, fmt.Sprintf("%s=$%d", "updated_at", paramIndex))
	args = append(args, time.Now())
	paramIndex++

	str := strings.Join(updates, ", ")
	whereClause := fmt.Sprintf("WHERE id = $%d RETURNING id, user_id, ip_address, device_info, is_active, expired_at, updated_at, created_at", paramIndex)
	args = append(args, s.sessionId)

	return baseQuery + " " + str + " " + whereClause, args, nil
}
