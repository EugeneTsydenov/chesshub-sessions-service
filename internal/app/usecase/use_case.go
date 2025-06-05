package usecase

import "context"

type UseCase[Input comparable, Output comparable] interface {
	Execute(context.Context, Input) (Output, error)
}
