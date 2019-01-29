package status

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ErrInternal represents internal server error.
var ErrInternal = ErrServiceStatus{
	ServiceStatus{Code: http.StatusInternalServerError, Message: "Internal Server Error"},
}

// ErrNotFound represents an error when a domain artifact was not found.
var ErrNotFound = ErrServiceStatus{
	ServiceStatus{Code: http.StatusNotFound, Message: "Not Found"},
}

// ErrBadRequest represents an invalid request error.
var ErrBadRequest = ErrServiceStatus{
	ServiceStatus{Code: http.StatusBadRequest, Message: "Bad Request"},
}

// ErrUnauhtorized represents an unauthorized request error.
var ErrUnauhtorized = ErrServiceStatus{
	ServiceStatus{Code: http.StatusUnauthorized, Message: "Unauthorized"},
}

// ErrNotImplemented represents an unauthorized request error.
var ErrNotImplemented = ErrServiceStatus{
	ServiceStatus{Code: http.StatusNotImplemented, Message: "Not Implemented"},
}

// ErrContentTypeNotSupported represents unsupported media type.
var ErrContentTypeNotSupported = ErrServiceStatus{
	ServiceStatus{Code: http.StatusUnsupportedMediaType, Message: "Unsupported Media Type"},
}

// ErrStatusConflict represents conflict because of inconsistent or duplicated info.
var ErrStatusConflict = ErrServiceStatus{
	ServiceStatus{Code: http.StatusConflict, Message: "Conflict because of inconsistent or duplicated info"},
}

// Success represents a generic success.
var Success = ServiceStatus{Code: http.StatusOK, Message: "OK"}

// ServiceStatus captures basic information about a status construct.
type ServiceStatus struct {
	Code    int          `json:"code,omitempty"`
	Message string       `json:"msg"`
	Details []*StatusDtl `json:"details,omitempty"`
}

// StatusDtl captures basic information about a status construct.
type StatusDtl struct {
	Code    string `json:"-"`
	Message string `json:"msg"`
}

// ErrServiceStatus captures basic information about an error.
type ErrServiceStatus struct {
	ServiceStatus
}

// WithMessage returns an error status with given message.
func (e ErrServiceStatus) WithMessage(msg string) ErrServiceStatus {
	return ErrServiceStatus{ServiceStatus{Code: e.Code, Message: strings.Join([]string{e.Message, msg}, ": ")}}
}

// WithError returns an error status with given err.Error().
func (e ErrServiceStatus) WithError(err error) ErrServiceStatus {
	return ErrServiceStatus{ServiceStatus{Code: e.Code, Message: strings.Join([]string{e.Message, err.Error()}, ": ")}}
}

// AddDtlMsg returns an error status with given message.
func (e *ErrServiceStatus) AddDtlMsg(msgs ...string) {
	if e.Details == nil {
		e.Details = []*StatusDtl{}
	}
	for _, m := range msgs {
		d := &StatusDtl{Message: m}
		e.Details = append(e.Details, d)
	}
}

// AddDtl returns an error status with given message.
func (e *ErrServiceStatus) AddDtl(code, msg string) {
	if e.Details == nil {
		e.Details = []*StatusDtl{}
	}
	d := &StatusDtl{Code: code, Message: msg}
	e.Details = append(e.Details, d)
}

// New returns a new status with given status instance.
func New(ss ServiceStatus) ServiceStatus {
	return ServiceStatus{ss.Code, ss.Message, nil}
}

// NewError returns a new status with given status instance.
func NewErrorStatus(err ErrServiceStatus) ServiceStatus {
	return ServiceStatus{err.Code, err.Message, nil}
}

// NewUserDefined returns a new status with given code and message.
func NewUserDefined(code int, msg string) ServiceStatus {
	return ServiceStatus{Code: code, Message: msg}
}

func (e ErrServiceStatus) Error() string {
	if errB, err := json.Marshal(&e); err == nil {
		return string(errB)
	}
	return `{"code":500, "msg": "error marshal failed"}`
}
