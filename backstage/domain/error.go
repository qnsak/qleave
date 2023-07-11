package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidatorErrorResponse struct {
	Message map[string][]string `json:"message"`
}
