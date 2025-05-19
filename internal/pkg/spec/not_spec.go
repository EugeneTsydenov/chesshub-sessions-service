package spec

import "fmt"

type NotSpec struct {
	Spec Spec
}

func NewNot(spec Spec) *NotSpec {
	return &NotSpec{Spec: spec}
}

func (s *NotSpec) ToSQL() (string, []any) {
	sqlPart, args := s.Spec.ToSQL()
	return fmt.Sprintf("NOT (%s)", sqlPart), args
}
