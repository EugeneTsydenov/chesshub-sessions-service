package grpcerrors

import (
	"errors"
	"fmt"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToGRPCError(err error) error {
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		return appErrorToGRPCError(appErr)
	}

	return status.Error(codes.Unknown, err.Error())
}

func appErrorToGRPCError(err *apperrors.AppError) error {
	switch err.Code {
	case apperrors.InvalidArgument:
		return withDetails(codes.InvalidArgument, err)
	case apperrors.NotFound:
		return status.Error(codes.NotFound, err.Error())
	case apperrors.Conflict:
		return withDetails(codes.AlreadyExists, err)
	case apperrors.Internal:
		return status.Error(codes.Internal, err.Error())
	case apperrors.Unauthenticated:
		return status.Error(codes.Unauthenticated, err.Error())
	case apperrors.Forbidden:
		return status.Error(codes.PermissionDenied, err.Error())
	case apperrors.Canceled:
		return status.Error(codes.Canceled, err.Error())
	case apperrors.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}

func withDetails(code codes.Code, err *apperrors.AppError) error {
	errInfo := &errdetails.ErrorInfo{
		Reason:   err.Code.String(),
		Domain:   "user",
		Metadata: err.Details,
	}

	st := status.New(code, err.Message)
	detailedStatus, detailErr := st.WithDetails(errInfo)
	if detailErr != nil {
		return fmt.Errorf("st.WithDetails: %w", detailErr)
	}
	return detailedStatus.Err()
}
