package sessionspec

import (
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/pkg/spec"
)

type SessionFilterSpec struct {
	filterMap map[string]string
}

func NewSessionFilterSpec(filterMap map[string]string) *SessionFilterSpec {
	return &SessionFilterSpec{
		filterMap: filterMap,
	}
}

func (s *SessionFilterSpec) ToSQL() (string, []any, error) {
	baseQuery := "SELECT * FROM sessions"
	var specs []spec.Spec
	paramIndex := 1

	for key, value := range s.filterMap {
		column, ok := fieldToColumn[key]
		if !ok {
			return "", nil, errors.NewInvalidFieldError(fmt.Sprintf("invalid filter field: '%s' is not recognized", key))
		}

		specs = append(specs, spec.NewFieldSpec(column, "=", paramIndex, value))
		paramIndex++
	}

	andSpec := spec.NewAnd(specs...)

	query, args, err := andSpec.ToSQL()
	if err != nil {
		return "", nil, err
	}

	return baseQuery + " WHERE " + query, args, nil
}
