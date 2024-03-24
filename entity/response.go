package entity

type Response struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Employee any    `json:"data,omitempty"`
}
