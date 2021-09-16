package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ForbiddenResponse struct {
	Error baseError `json:"error"`
}

func NewForbiddenResponse() *ForbiddenResponse {
	r := ForbiddenResponse{}
	r.Error.Code = http.StatusForbidden
	b := Forbidden
	r.Error.Message = &b
	return &r
}

func (r ForbiddenResponse) GetStatusCode() int {
	return r.Error.Code
}

func (r *ForbiddenResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func (r *ForbiddenResponse) SetCode(code int) *ForbiddenResponse {
	r.Error.Code = code
	return r
}

func (r *ForbiddenResponse) SetErrors(errors []string) *ForbiddenResponse {
	r.Error.Errors = errors
	return r
}

func (r *ForbiddenResponse) SetMessage(message string) *ForbiddenResponse {
	r.Error.Message = &message
	return r
}
