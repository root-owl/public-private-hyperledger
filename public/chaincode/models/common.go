package models

import (
	"time"
)

type MetaInfo struct {
	DocType   DocType      `json:"doc_type"`
	CreatedAt time.Time    `json:"created_at"`
	Status    MemberStatus `json:"status"`
}
