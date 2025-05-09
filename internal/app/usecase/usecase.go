package usecase

import "context"

type Executor[Input comparable, Output comparable] interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
