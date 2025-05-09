package repo

import "context"

type creator[Input comparable, Output comparable] interface {
	Create(ctx context.Context, input Input) (Output, error)
}

type reader[Output comparable] interface {
	Read(ctx context.Context, spec any) (Output, error)
}

type updater[Input comparable, Output comparable] interface {
	Update(ctx context.Context, input Input) (Output, error)
}

type deleter[Input comparable, Output comparable] interface {
	Delete(ctx context.Context, input Input) (Output, error)
}

type repo[Model comparable] interface {
	creator[Model, Model]
	reader[Model]
	updater[Model, Model]
	deleter[Model, Model]
}
