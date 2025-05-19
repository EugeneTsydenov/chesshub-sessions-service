package spec

import "fmt"

type NotSpec struct {
	Spec Spec
}

func NewNot(spec Spec) *NotSpec {
	return &NotSpec{Spec: spec}
}

func (s *NotSpec) ToSQL() (string, []any, error) {
	sqlPart, args, err := s.Spec.ToSQL()
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf("NOT (%s)", sqlPart), args, nil
}
