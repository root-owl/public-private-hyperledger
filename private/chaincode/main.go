package main

import (
	"fmt"

	"secc/private/chaincode/models"

	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/response"
	"github.com/s7techlab/cckit/router/param"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/s7techlab/cckit/router"
)

type SmartContract struct {
	router *router.Group
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return response.Create(owner.SetFromCreator(t.router.Context(stub)))
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return t.router.Handle(stub)
}

func NewRouter() *SmartContract {

	router := router.New("SECC")
	contract := &SmartContract{router: router}

	router.Group(`User`).
		Invoke(`Register`, contract.registerUser, param.Struct(`request`, &models.LocalUserPayload{})).
		Invoke(`Login`, contract.loginUser, param.Struct(`request`, &models.LoginRequest{})).
		Invoke(`ChangeAccessStatus`, contract.changeAccessStatus, param.Struct(`request`, &models.AccessStatusUpdatePayload{}), owner.Only)

	router.Group(`AppServer`).
		Invoke(`Register`, contract.registerApplicationServer, param.Struct(`request`, &models.AppServerPayload{}))

	router.Group(`HIDS`).
		Invoke(`Scan`, contract.submitHIDSScan, param.Struct(`request`, &models.HIDSScanPayload{})).
		Invoke(`IncidentReport`, contract.submitHIDSIncedentReport, param.Struct(`request`, &models.HIDSIncidentReportPayload{})).
		Invoke(`Whitelist`, contract.submitHIDSWhitelistReport, param.Struct(`request`, &models.HIDSWhitelistPayload{})).
		Invoke(`MasterTree`, contract.submitHIDSMasterTree, param.Struct(`request`, &models.HIDSMasterTreePayload{})).
		Query(`GetServerIncidents`, contract.getServerAndIncidents, param.String(`id`))

	return contract
}

func main() {

	if err := shim.Start(NewRouter()); err != nil {
		fmt.Println("Error starting chaincode")
	}

}
