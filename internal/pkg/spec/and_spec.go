package spec

import "fmt"

type AndSpec struct {
	Specs []Spec
}

func NewAnd(specs ...Spec) *AndSpec {
	return &AndSpec{Specs: specs}
}

func (s *AndSpec) ToSQL() (string, []any) {
	var sqlParts []string
	var args []any

	for _, spec := range s.Specs {
		sqlPart, specArgs := spec.ToSQL()
		sqlParts = append(sqlParts, fmt.Sprintf("(%s)", sqlPart))
		args = append(args, specArgs...)
	}

	return joinWithOperator(sqlParts, "AND"), args
}
