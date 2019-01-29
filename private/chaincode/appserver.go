package main

import (
	"encoding/json"
	"fmt"
	"secc/private/chaincode/core/status"
	"secc/private/chaincode/models"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"github.com/s7techlab/cckit/router"
)

func (contract *SmartContract) registerApplicationServer(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.AppServerPayload)
	requestBody.AppServer.MetaInfo = models.MetaInfo{DocType: models.DocTypeAppServer,
		CreatedAt: time.Now()}

	if insErr := context.State().Insert(requestBody.ID, requestBody.AppServer); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(200, "Application server registered"), nil
}

func (contract *SmartContract) getServerAndIncidents(context router.Context) (interface{}, error) {
	context.Logger().SetLevel(shim.LogDebug)

	userID := context.Arg(`id`).(string)

	localUser, queryErr := context.State().Get(userID, &models.LocalUser{})

	if queryErr != nil {
		return nil, status.ErrNotFound.WithError(queryErr)
	}

	context.Logger().Debug(localUser)

	user := localUser.(models.LocalUser)

	if user.Status != models.Approved {
		return nil, status.ErrUnauhtorized.WithMessage("You are not authorised to access this functionality")
	}

	serverQuery := fmt.Sprintf("{" +
		"\"selector\":{" +
		"\"meta_info.doc_type\":\"application_server\"" +
		"}" +
		"}")

	queryIteratorInterface, queryIteratorErr := context.Stub().GetQueryResult(serverQuery)

	if queryIteratorErr != nil {
		return nil, status.ErrInternal.WithError(queryIteratorErr)
	}

	appServersPayloads := make([]models.ServerIncidentPayload, 0, 0)

	for queryIteratorInterface.HasNext() {
		queryKV, queryerr := queryIteratorInterface.Next()

		if queryerr != nil {
			return nil, status.ErrInternal.WithError(queryerr)
		}

		var appServer models.AppServer
		unmarshalErr := json.Unmarshal(queryKV.Value, &appServer)

		if unmarshalErr != nil {
			return nil, status.ErrInternal.WithError(unmarshalErr)
		}

		incidentQuery := fmt.Sprintf("{" +
			"\"selector\":{" +
			"\"meta_info.doc_type\":\"hids_incident_report\"," +
			"\"server_id\":" + "\"" + queryKV.Key + "\"" +
			"}" +
			"}")
		context.Logger().Error(incidentQuery)
		incidentQueryIteratorInterface, incidentQueryIteratorErr := context.Stub().GetQueryResult(incidentQuery)

		if incidentQueryIteratorErr != nil {
			return nil, status.ErrInternal.WithError(incidentQueryIteratorErr)
		}

		incidents := make([]models.HIDSIncidentReportPayload, 0, 0)

		for incidentQueryIteratorInterface.HasNext() {
			incidentQueryKV, KVError := incidentQueryIteratorInterface.Next()

			if KVError != nil {
				return nil, status.ErrInternal.WithError(KVError)
			}

			var hidsIncidentReport models.HIDSIncidentReport

			incidentRpUnmarshalErr := json.Unmarshal(incidentQueryKV.Value, &hidsIncidentReport)

			if incidentRpUnmarshalErr != nil {
				return nil, status.ErrInternal.WithError(incidentRpUnmarshalErr)
			}

			incidentPayload := models.HIDSIncidentReportPayload{
				ID:                 incidentQueryKV.Key,
				HIDSIncidentReport: hidsIncidentReport}

			incidents = append(incidents, incidentPayload)
		}

		server := models.AppServerPayload{
			ID:        queryKV.Key,
			AppServer: appServer}

		incidentPayload := models.ServerIncidentPayload{
			AppServerPayload:          server,
			HIDSIncidentReportPayload: incidents}

		appServersPayloads = append(appServersPayloads, incidentPayload)
	}

	serverIncidents := models.AllServerIncidentsPayload{
		Data: appServersPayloads}

	return serverIncidents, nil
}
