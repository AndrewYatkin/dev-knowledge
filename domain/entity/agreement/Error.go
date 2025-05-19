package agreementEntity

import "dev-knowledge/infrastructure/errors"

var (
	ErrAcceptedDateIsRequired        = errors.NewError("5c7f9619-001", "Accepted date is required")
	ErrUnsupportedDateForNotAccepted = errors.NewError("5c7f9619-002", "Accepted date not supported for not accepted agreement")
)
