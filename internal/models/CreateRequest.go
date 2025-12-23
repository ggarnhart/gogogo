package models

type CreateRequest struct {
	Message   string `json:"message"`
	IPAddress string `json:"ip_address"`
}
