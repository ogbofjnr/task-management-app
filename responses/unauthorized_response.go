package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UnauthorizedResponse struct {
	Error baseError `json:"error"`
}

func NewUnauthorizedResponse() *UnauthorizedResponse {
	r := UnauthorizedResponse{}
	r.Error.Code = http.StatusUnauthorized
	b := Unauthorized
	r.Error.Message = &b
	return &r
}

func (r *UnauthorizedResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func (r *UnauthorizedResponse) GetStatusCode() int {
	return r.Error.Code
}

func (r *UnauthorizedResponse) SetCode(code int) *UnauthorizedResponse {
	r.Error.Code = code
	return r
}

func (r *UnauthorizedResponse) SetErrors(errors []string) *UnauthorizedResponse {
	r.Error.Errors = errors
	return r
}

func (r *UnauthorizedResponse) SetMessage(message string) *UnauthorizedResponse {
	r.Error.Message = &message
	return r
}
