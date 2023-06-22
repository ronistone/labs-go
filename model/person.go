package model

import "time"

type Person struct {
	ID        *int64     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
