package grpcerrors

import (
	"errors"
	"fmt"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/apperrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapAppErrorToGrpcError(err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		grpcCode := mapAppErrorToGrpcCode(appErr.Code)

		errInfo := &errdetails.ErrorInfo{
			Reason:   appErr.Code.String(),
			Domain:   "user",
			Metadata: appErr.Details,
		}

		st := status.New(grpcCode, appErr.Message)
		detailedStatus, detailErr := st.WithDetails(errInfo)
		if detailErr != nil {
			return fmt.Errorf("st.WithDetails: %w", detailErr)
		}
		return detailedStatus.Err()
	}

	errInfo := &errdetails.ErrorInfo{
		Reason:   "UNKNOWN_ERROR",
		Domain:   "user",
		Metadata: map[string]string{},
	}

	st := status.New(codes.Unknown, err.Error())
	detailedStatus, detailErr := st.WithDetails(errInfo)
	if detailErr != nil {
		return fmt.Errorf("st.WithDetails: %w", detailErr)
	}
	return detailedStatus.Err()
}

func mapAppErrorToGrpcCode(c apperrors.ErrorCode) codes.Code {
	switch c {
	case apperrors.Conflict:
		return codes.AlreadyExists
	case apperrors.Internal:
		return codes.Internal
	case apperrors.Unauthenticated:
		return codes.Unauthenticated
	case apperrors.Forbidden:
		return codes.PermissionDenied
	case apperrors.NotFound:
		return codes.NotFound
	case apperrors.InvalidArgument:
		return codes.InvalidArgument
	case apperrors.DeadlineExceeded:
		return codes.DeadlineExceeded
	case apperrors.Canceled:
		return codes.Canceled
	default:
		return codes.Unknown
	}
}
