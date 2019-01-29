package models

type HIDSScan struct {
	ServerID          string   `json:"server_id,omitempty"`
	Version           string   `json:"version,omitempty"`
	Parameters        string   `json:"parameters,omitempty"`
	MasterTreeRef     string   `json:"master_tree_ref,omitempty"`
	JobRef            string   `json:"job_ref,omitempty"`
	WhitelistRef      string   `json:"whitelist_ref,omitempty"`
	ScanResultHash    string   `json:"scan_result_hash,omitempty"`
	ServerSignature   string   `json:"server_signature,omitempty"`
	Status            bool     `json:"status"`
	IncidentReportRef string   `json:"incident_report_ref,omitempty"`
	AdminRef          string   `json:"admin_reference,omitempty"`
	MetaInfo          MetaInfo `json:"meta_info,omitempty"`
}

type HIDSIncidentReport struct {
	ServerID       string   `json:"server_id,omitempty"`
	ScanID         string   `json:"scan_id,omitempty"`
	Description    string   `json:"description,omitempty"`
	DeletedFiles   string   `json:"deleted_files,omitempty"`
	ModifiedFiles  string   `json:"modified_files,omitempty"`
	CreatedFiles   string   `json:"created_files,omitempty"`
	Replay         string   `json:"replay,omitempty"`
	MissedScanInfo string   `json:"missed_scan_info,omitempty"`
	EmailRef       string   `json:"email_ref,omitempty"`
	MetaInfo       MetaInfo `json:"meta_info,omitempty"`
}

type HIDSWhitelist struct {
	ServerID  string   `json:"server_id,omitempty"`
	Whitelist string   `json:"whitelist,omitempty"`
	AdminRef  string   `json:"admin_ref,omitempty"`
	Version   string   `json:"version,omitempty"`
	MetaInfo  MetaInfo `json:"meta_info,omitempty"`
}

type HIDSMasterTree struct {
	ServerID       string   `json:"server_id,omitempty"`
	MasterTree     string   `json:"master_tree,omitempty"`
	MasterTreeHash string   `json:"master_tree_hash,omitempty"`
	AdminRef       string   `json:"admin_ref,omitempty"`
	WhitelistRef   string   `json:"whitelist_ref,omitempty"`
	ScanIDRef      string   `json:"scan_id_ref,omitempty"`
	MetaInfo       MetaInfo `json:"meta_info,omitempty"`
}
