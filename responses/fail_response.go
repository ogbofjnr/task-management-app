package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InternalServerErrorResponse struct {
	Error baseError `json:"error"`
}

func NewInternalServerErrorResponse() *InternalServerErrorResponse {
	r := InternalServerErrorResponse{}
	r.Error.Code = http.StatusInternalServerError
	return &r
}

func (r *InternalServerErrorResponse) GetStatusCode() int {
	return r.Error.Code
}

func (r *InternalServerErrorResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}
