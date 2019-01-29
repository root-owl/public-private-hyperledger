package models

type AppServer struct {
	IPAddress   string   `json:"ip_address,omitempty"`
	AdminRef    string   `json:"admin_ref,omitempty"`
	Certificate string   `json:"cert,omitempty"`
	MetaInfo    MetaInfo `json:"meta_info,omitempty"`
}
