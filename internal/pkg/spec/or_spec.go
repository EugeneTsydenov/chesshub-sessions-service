package spec

import "fmt"

type OrSpec struct {
	Specs []Spec
}

func NewOr(specs ...Spec) *OrSpec {
	return &OrSpec{Specs: specs}
}

func (s *OrSpec) ToSQL() (string, []any) {
	var sqlParts []string
	var args []any

	for _, spec := range s.Specs {
		sqlPart, specArgs := spec.ToSQL()
		sqlParts = append(sqlParts, fmt.Sprintf("(%s)", sqlPart))
		args = append(args, specArgs...)
	}

	return joinWithOperator(sqlParts, "OR"), args
}
