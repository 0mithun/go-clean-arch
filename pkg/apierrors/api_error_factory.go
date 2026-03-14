package apierrors

import (
	"errors"
	"net/http"
)

func NewAPIError(statusCode int, msg string) APIError {

	return &apiError{
		XMessage:    msg,
		XStatusCode: statusCode,
	}
}

func NewNotFoundError(msg string) APIError {
	return NewAPIError(http.StatusNotFound, msg)
}

func NewBadRequestError(msg string) APIError {
	return NewAPIError(http.StatusBadRequest, msg)
}
func NewInternalServerError(msg string) APIError {
	return NewAPIError(http.StatusInternalServerError, msg)
}

func NewUnauthorizedError(msg string) APIError {
	return NewAPIError(http.StatusUnauthorized, msg)
}

func NewForbiddenError(msg string) APIError {
	return NewAPIError(http.StatusForbidden, msg)
}

func NewUnimplementedError(msg string) APIError {
	return NewAPIError(http.StatusNotImplemented, msg)
}

func FromError(err error) APIError {
	if err == nil {
		return nil
	}

	var apiErr APIError
	if !errors.As(err, &apiErr) {
		return NewAPIError(http.StatusInternalServerError, err.Error())
	}

	return apiErr
}
