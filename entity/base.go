package entity

import "time"

// BaseFilter is generic filter will be used by other complex filter
type BaseFilter struct {
	Num    int32
	Cursor string
}

// BaseEntity is generic entity for other entity
// Status present status of row
// default is 1 // active
// 0 is deleted
type BaseEntity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    int8      `json:"status"`
}
