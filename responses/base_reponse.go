package responses

type Response struct {
	StatusCode int `json:"-"`
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}
