package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResponse struct {
	Code    int     `json:"code"`
	Message *string `json:"message"`
}

func NewSuccessResponse() *SuccessResponse {
	r := &SuccessResponse{}
	r.Code = http.StatusOK
	b := Success
	r.Message = &b
	return r
}

func (r *SuccessResponse) GetStatusCode() int {
	return r.Code
}

func (r *SuccessResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func (r *SuccessResponse) SetCode(code int) *SuccessResponse {
	r.Code = code
	return r
}

func (r *SuccessResponse) SetMessage(message string) *SuccessResponse {
	r.Message = &message
	return r
}
