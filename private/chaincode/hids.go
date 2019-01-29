package main

import (
	"fmt"
	"secc/private/chaincode/core/status"
	"secc/private/chaincode/models"
	"time"

	"github.com/s7techlab/cckit/router"
)

//Submit HIDS Scan
func (contract *SmartContract) submitHIDSScan(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.HIDSScanPayload)
	requestBody.HIDSScan.MetaInfo = models.MetaInfo{DocType: models.DocTypeHIDSScan,
		CreatedAt: time.Now()}

	if exist := IdExist(requestBody.HIDSScan.ServerID, context); !exist {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Server: %s does not exist", requestBody.HIDSScan.ServerID))
	}

	if insErr := context.State().Insert(requestBody.ID, requestBody.HIDSScan); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(200, "HIDS scan saved successfully"), nil
}

//Submit HIDS Incident Report
func (contract *SmartContract) submitHIDSIncedentReport(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.HIDSIncidentReportPayload)
	requestBody.HIDSIncidentReport.MetaInfo = models.MetaInfo{DocType: models.DocTypeHIDSIncidentReport,
		CreatedAt: time.Now()}

	if exist := IdExist(requestBody.HIDSIncidentReport.ServerID, context); !exist {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Server: %s does not exist", requestBody.HIDSIncidentReport.ServerID))
	} else if exist := IdExist(requestBody.HIDSIncidentReport.ScanID, context); !exist {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("HIDS Scan: %s does not exist", requestBody.HIDSIncidentReport.ScanID))
	}

	if insErr := context.State().Insert(requestBody.ID, requestBody.HIDSIncidentReport); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(200, "HIDS incident report submitted successfully"), nil
}

//Submit HIDS Whitelist
func (contract *SmartContract) submitHIDSWhitelistReport(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.HIDSWhitelistPayload)
	requestBody.HIDSWhitelist.MetaInfo = models.MetaInfo{DocType: models.DocTypeHIDSWhitelist,
		CreatedAt: time.Now()}

	if exist := IdExist(requestBody.HIDSWhitelist.ServerID, context); !exist {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Server: %s does not exist", requestBody.HIDSWhitelist.ServerID))
	}

	if insErr := context.State().Insert(requestBody.ID, requestBody.HIDSWhitelist); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(200, "HIDS whitelist submitted successfully"), nil
}

//Submit HIDS Master Tree
func (contract *SmartContract) submitHIDSMasterTree(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.HIDSMasterTreePayload)
	requestBody.HIDSMasterTree.MetaInfo = models.MetaInfo{DocType: models.DocTypeHIDSMasterTree,
		CreatedAt: time.Now()}

	if exist := IdExist(requestBody.HIDSMasterTree.ServerID, context); !exist {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Server: %s does not exist", requestBody.HIDSMasterTree.ServerID))
	}

	if insErr := context.State().Insert(requestBody.ID, requestBody.HIDSMasterTree); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(200, "HIDS whitelist submitted successfully"), nil
}
