package models

type MemberRole int

const (
	Admin MemberRole = 1
	User  MemberRole = 0
)

type AccessStatus int

const (
	NoAccess     AccessStatus = 0
	Approved     AccessStatus = 1
	DeniedAccess AccessStatus = 2
)

type DocType string

const (
	DocTypeLocalUser          DocType = "local_user"
	DocTypeAppServer          DocType = "application_server"
	DocTypeHIDSScan           DocType = "hids_scan"
	DocTypeHIDSIncidentReport DocType = "hids_incident_report"
	DocTypeHIDSWhitelist      DocType = "hids_whitelist"
	DocTypeHIDSMasterTree     DocType = "hids_master_tree"
)
