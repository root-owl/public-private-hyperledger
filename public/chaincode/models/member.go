package models

type Member struct {
	ID            string     `json:"member_id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	PublicKey     string     `json:"public_key"`
	PublicAddress string     `json:"public_address"`
	Description   string     `json:"description"`
	EmailAddress  string     `json:"email_address"`
	Password      string     `json:"password"`
	Role          MemberRole `json:"role"`
	MetaInfo      MetaInfo   `json:"meta_info"`
}

type AccountModel struct {
	ID           string   `json:"account_id"`
	OwnerID      string   `json:"owner_id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Balance      float64  `json:"balance"`
	Subscription []string `json:"subscription"`
}
