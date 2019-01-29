package models

type MemberRole int

const (
	Admin MemberRole = 1
	User  MemberRole = 2
)

type DocType string

const (
	DocTypeMember DocType = "member"
)

type MemberStatus int

const (
	Active             MemberStatus = 1
	Inactive           MemberStatus = -1
	PendingForApproval MemberStatus = 0
)
