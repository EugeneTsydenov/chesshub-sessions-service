package spec

type Spec interface {
	BuildQuery() (string, []any)
}
