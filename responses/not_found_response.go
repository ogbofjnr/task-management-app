package responses

import (
	"encoding/json"
	"fmt"
	"github.com/ogbofjnr/maze/utils"
	"net/http"
)

type NotFoundResponse struct {
	Error baseError `json:"error"`
}

func NewNotFoundRequest() *NotFoundResponse {
	r := NotFoundResponse{}
	r.Error.Code = http.StatusNotFound
	r.Error.Message = utils.StrToPointer(NotFound)
	return &r
}

func (r NotFoundResponse) GetStatusCode() int {
	return r.Error.Code
}

func (r *NotFoundResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func (r *NotFoundResponse) SetCode(code int) *NotFoundResponse {
	r.Error.Code = code
	return r
}

func (r *NotFoundResponse) SetErrors(errors []string) *NotFoundResponse {
	r.Error.Errors = errors
	return r
}

func (r *NotFoundResponse) SetMessage(message string) *NotFoundResponse {
	r.Error.Message = &message
	return r
}
