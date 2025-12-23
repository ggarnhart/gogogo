package models

import "time"

type Request struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}
