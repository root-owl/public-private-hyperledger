package main

import (
	"encoding/json"
	"fmt"
	"time"

	"secc/private/chaincode/core/status"
	"secc/private/chaincode/models"

	"github.com/s7techlab/cckit/router"
	"github.com/shomali11/util/xhashes"
)

//Register user
func (contract *SmartContract) registerUser(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.LocalUserPayload)
	requestBody.LocalUser.MetaInfo = models.MetaInfo{DocType: models.DocTypeLocalUser,
		CreatedAt: time.Now()}

	requestBody.LocalUser.Password = xhashes.SHA256(requestBody.LocalUser.Password)

	if insErr := context.State().Insert(requestBody.ID, requestBody.LocalUser); insErr != nil {
		return nil, status.ErrInternal.WithError(insErr)
	}

	return status.NewUserDefined(201, "User registered successfully"), nil
}

//Login User
func (contract *SmartContract) loginUser(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.LoginRequest)

	hashedPassword := xhashes.SHA256(requestBody.Password)

	userQuery := fmt.Sprintf("{"+
		"\"selector\":{"+
		"\"meta_info.doc_type\":\"%s\","+
		"\"email\":\"%s\","+
		"\"password\":\"%s\""+
		"}}", models.DocTypeLocalUser, requestBody.Email, hashedPassword)

	context.Logger().Errorf("Query ==> %s", userQuery)

	queryIterator, queryErr := context.Stub().GetQueryResult(userQuery)

	if queryErr != nil {
		return nil, status.ErrInternal.WithError(queryErr)
	}

	if queryIterator.HasNext() {
		queryKV, err := queryIterator.Next()
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		var member models.LocalUser
		unMarshallErr := json.Unmarshal(queryKV.Value, &member)
		if err != nil {
			return nil, status.ErrInternal.WithError(unMarshallErr)
		}

		member.Password = ""

		localUserPayload := models.LocalUserPayload{
			ID:        queryKV.Key,
			LocalUser: member}

		return localUserPayload, nil

	}

	return status.ErrNotFound.WithMessage("Please enter valid email and password"), nil
}

func (contract *SmartContract) changeAccessStatus(context router.Context) (interface{}, error) {

	requestBody := context.Arg(`request`).(models.AccessStatusUpdatePayload)

	//check if user id exist
	result, queryErr := context.State().Get(requestBody.UserID, &models.LocalUser{})

	if queryErr != nil {
		return nil, status.ErrNotFound.WithError(queryErr)
	}

	if requestBody.AccessStatus < 0 && requestBody.AccessStatus > 2 {
		return nil, status.ErrBadRequest.WithMessage("Invalid access status received")
	}

	user := result.(models.LocalUser)
	user.Status = requestBody.AccessStatus

	//save changes
	if putErr := context.State().Put(requestBody.UserID, user); putErr != nil {
		return nil, status.ErrNotFound.WithError(putErr)
	}

	return struct {
		Message string `json:"message"`
	}{
		Message: "Success"}, nil

}
