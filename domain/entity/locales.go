package entity

import "time"

type Locale struct {
	LocalID   int        `json:"local_id"`
	LoaclCode string     `json:"local_code"`
	LocalName string     `json:"local_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Disabed   bool       `json:"disabed"`
}
