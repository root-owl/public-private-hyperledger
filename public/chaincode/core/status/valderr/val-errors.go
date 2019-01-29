package valderr

import (
	"secc/public/chaincode/core/status"

	validation "github.com/go-ozzo/ozzo-validation"
)

func NewErrStatusWithValErrors(e status.ErrServiceStatus, valErrs validation.Errors) status.ErrServiceStatus {
	errSvc := status.ErrServiceStatus{ServiceStatus: status.ServiceStatus{Code: e.Code, Message: e.Message, Details: nil}}
	for _, msg := range valErrs {
		errSvc.AddDtlMsg(msg.Error())
	}
	return errSvc
}
