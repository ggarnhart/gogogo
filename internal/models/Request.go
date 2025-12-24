package models

import (
	"net/netip"
	"time"
)

type Request struct {
	ID        string     `db:"id" json:"id"`
	Message   string     `db:"message" json:"message"`
	IPAddress netip.Addr `db:"ip_address" json:"ip_address"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	Status    string     `db:"status" json:"status"`
}
