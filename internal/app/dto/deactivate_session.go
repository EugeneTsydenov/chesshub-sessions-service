package dto

type (
	DeactivateSessionInputDTO struct {
		SessionID string
	}

	DeactivateSessionOutputDTO struct {
		Success bool
		Message string
	}
)
