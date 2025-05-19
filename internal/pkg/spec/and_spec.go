package spec

import "fmt"

type AndSpec struct {
	Specs []Spec
}

func NewAnd(specs ...Spec) *AndSpec {
	return &AndSpec{Specs: specs}
}

func (s *AndSpec) ToSQL() (string, []any, error) {
	var sqlParts []string
	var args []any

	for _, spec := range s.Specs {
		sqlPart, specArgs, err := spec.ToSQL()
		if err != nil {
			return "", nil, err
		}

		sqlParts = append(sqlParts, fmt.Sprintf("(%s)", sqlPart))
		args = append(args, specArgs...)
	}

	return joinWithOperator(sqlParts, "AND"), args, nil
}
