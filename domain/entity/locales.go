package entity

import "time"

type Locale struct {
	LocaleID   int        `json:"locale_id"`
	LocaleCode string     `json:"locale_code"`
	LocaleName string     `json:"local_name"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Disabed    bool       `json:"disabed"`
}
