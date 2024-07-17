package payload

type Response struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Error   any               `json:"error,omitempty"`
	Errors  []ErrorValidation `json:"errors,omitempty"`
	Data    any               `json:"data,omitempty"`
}

type ErrorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
