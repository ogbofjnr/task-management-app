package responses

type baseError struct {
	Code    int      `json:"code"`
	Errors  []string `json:"errors"`
	Message *string  `json:"message"`
}
