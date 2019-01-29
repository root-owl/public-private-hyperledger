package models

type LocalUser struct {
	FirstName   string       `json:"first_name,omitempty"`
	LastName    string       `json:"last_name,omitempty"`
	Email       string       `json:"email,omitempty"`
	Password    string       `json:"password,omitempty"`
	JobTitle    string       `json:"job_title,omitempty"`
	Role        MemberRole   `json:"member_role"`
	Status      AccessStatus `json:"access_status"`
	Certificate string       `json:"cert,omitempty"`
	MetaInfo    MetaInfo     `json:"meta_info,omitempty"`
}
