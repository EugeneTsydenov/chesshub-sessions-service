package spec

import "fmt"

type Spec interface {
	ToSQL() (string, []any)
}

func joinWithOperator(parts []string, operator string) string {
	return join(parts, fmt.Sprintf(" %s ", operator))
}

func join(parts []string, sep string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += sep
		}
		result += part
	}
	return result
}
