package errors

import (
	"bff/internal/api/dto"
	"bff/log"
	goerrors "errors"

	"github.com/pkg/errors"
)

var (
	ErrNotFoundField = errors.New("field not found")
	ErrMock          = errors.New("mock-error")
)

type AppError struct {
	Code int
	Err  error
}

func HandleAPPError(err error) (dto.Response, error) {
	log.Logger.Println(err)

	var appError AppError
	if goerrors.Is(err, appError) {
		return ErrorResponse(appError.PrettyMsg(), appError.Err.Error()), nil
	}
	return ErrorResponse(err.Error()), nil
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func (e AppError) HTTPStatusCode() int {
	return httpStatusCode(e.Code)
}

func (e AppError) PrettyMsg() string {
	return prettyMsg(e.Code)
}

func New(format string, args ...any) error {
	return &AppError{
		Code: InternalCode,
		Err:  errors.Errorf(format, args...),
	}
}

func Wrap(err error, message string) error {
	code := InternalCode

	var e *AppError

	if goerrors.As(err, &e) {
		code = e.Code
	}

	return &AppError{
		Code: code,
		Err:  errors.Wrap(err, message),
	}
}

func NewBadRequestError(message string, args ...any) error {
	return &AppError{
		Code: BadRequestCode,
		Err:  errors.Errorf(message, args...),
	}
}

func WrapBadRequestError(err error, message string, args ...any) error {
	return &AppError{
		Code: BadRequestCode,
		Err:  errors.Wrapf(err, message, args...),
	}
}

func WrapNotFoundError(err error, message string, args ...any) error {
	return &AppError{
		Code: NotFoundCode,
		Err:  errors.Wrapf(err, message, args...),
	}
}

func WrapInternalServerError(err error, message string, args ...any) error {
	return &AppError{
		Code: InternalCode,
		Err:  errors.Wrapf(err, message, args...),
	}
}
