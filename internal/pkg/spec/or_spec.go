package spec

import "fmt"

type OrSpec struct {
	Specs []Spec
}

func NewOr(specs ...Spec) *OrSpec {
	return &OrSpec{Specs: specs}
}

func (s *OrSpec) ToSQL() (string, []any, error) {
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

	return joinWithOperator(sqlParts, "OR"), args, nil
}
