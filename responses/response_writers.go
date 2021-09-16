package responses

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type ResponseInterface interface {
	GetStatusCode() int
	ToString() ([]byte, error)
}

func WriteResponse(w http.ResponseWriter, response ResponseInterface, logger *zap.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.GetStatusCode())
	data, err := response.ToString()
	if err != nil {
		logger.Error(fmt.Errorf("error writing response:%w", err).Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		logger.Error(fmt.Errorf("error writing response:%w", err).Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
