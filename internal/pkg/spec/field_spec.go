package spec

import "fmt"

type FieldSpec struct {
	Field      string
	Op         string
	ParamIndex int
	Value      any
}

func NewFieldSpec(field, op string, paramIndex int, value any) *FieldSpec {
	return &FieldSpec{Field: field, Op: op, ParamIndex: paramIndex, Value: value}
}

func (s *FieldSpec) ToSQL() (string, []any, error) {
	return fmt.Sprintf("%s %s $%d", s.Field, s.Op, s.ParamIndex), []any{s.Value}, nil
}
