package models

//********************************************
//				START
//		Request/Response Payloads
//********************************************
type AppServerPayload struct {
	ID        string    `json:"_id,omitempty"`
	AppServer AppServer `json:"app_server,omitempty"`
}

type LocalUserPayload struct {
	ID        string    `json:"_id,omitempty"`
	LocalUser LocalUser `json:"local_user,omitempty"`
}

type HIDSScanPayload struct {
	ID       string   `json:"_id,omitempty"`
	HIDSScan HIDSScan `json:"hids_scan,omitempty"`
}

type HIDSIncidentReportPayload struct {
	ID                 string             `json:"_id,omitempty"`
	HIDSIncidentReport HIDSIncidentReport `json:"hids_incident_report,omitempty"`
}

type HIDSWhitelistPayload struct {
	ID            string        `json:"_id,omitempty"`
	HIDSWhitelist HIDSWhitelist `json:"hids_whitelist,omitempty"`
}

type HIDSMasterTreePayload struct {
	ID             string         `json:"_id,omitempty"`
	HIDSMasterTree HIDSMasterTree `json:"hids_master_tree,omitempty"`
}

//********************************************
//					END
//********************************************

//********************************************
//				START
//		Structure for Server incidents
//********************************************
type AllServerIncidentsPayload struct {
	Data []ServerIncidentPayload `json:"data,omitempty"`
}

type ServerIncidentPayload struct {
	AppServerPayload          AppServerPayload            `json:"server,omitempty"`
	HIDSIncidentReportPayload []HIDSIncidentReportPayload `json:"incidents,omitempty"`
}

//********************************************
//					END
//********************************************

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AccessStatusUpdatePayload struct {
	UserID       string       `json:"user_id,omitempty"`
	AccessStatus AccessStatus `json:"access_status,omitempty"`
}
