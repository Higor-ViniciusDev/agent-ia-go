package internal_error

import "google.golang.org/grpc/codes"

type InternalError struct {
	Message  string
	Err      string
	CodeGrpc codes.Code
}

func (ie *InternalError) Error() string {
	return ie.Message + ": " + ie.Err
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "not_found",
		CodeGrpc: codes.InvalidArgument,
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "internal_server_error",
		CodeGrpc: codes.Internal,
	}
}

func NewBadRequestError(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "bad_request",
		CodeGrpc: codes.Aborted,
	}
}

func NewManyRequestError(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "many_request",
		CodeGrpc: codes.ResourceExhausted,
	}
}

func NewUnauthorizedAccess(menssage string) *InternalError {
	return &InternalError{
		Message:  menssage,
		Err:      "unauthorized",
		CodeGrpc: codes.Unauthenticated,
	}
}

func NewUnauthorizedEmailAlreadyExists(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "unauthorized_email_already_exists",
		CodeGrpc: codes.Unauthenticated,
	}
}

func NewUnauthorizedEmailNotVerified(message string) *InternalError {
	return &InternalError{
		Message:  message,
		Err:      "unauthorized_email_not_verified",
		CodeGrpc: codes.Unauthenticated,
	}
}
