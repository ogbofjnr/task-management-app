package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BadRequestResponse struct {
	Error baseError `json:"error"`
}

func NewBadRequestResponse() *BadRequestResponse {
	r := BadRequestResponse{}
	r.Error.Code = http.StatusBadRequest
	b := BadRequest
	r.Error.Message = &b
	return &r
}

func (r *BadRequestResponse) GetStatusCode() int {
	return r.Error.Code
}

func (r *BadRequestResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func (r *BadRequestResponse) SetCode(code int) *BadRequestResponse {
	r.Error.Code = code
	return r
}

func (r *BadRequestResponse) SetErrors(errors []string) *BadRequestResponse {
	r.Error.Errors = errors
	return r
}

func (r *BadRequestResponse) SetMessage(message string) *BadRequestResponse {
	r.Error.Message = &message
	return r
}
