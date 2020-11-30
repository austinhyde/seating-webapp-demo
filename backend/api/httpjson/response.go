package httpjson

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type HandlerFunc func(r *http.Request) Response

func WrapHandler(log *zerolog.Logger, h HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := h(r)
		w.WriteHeader(resp.Status)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Error().Err(err).Msg("could not json-encode response")
		}
	})
}

type Response struct {
	Status      int         `json:"status"`
	Success     bool        `json:"success"`
	ErrorMsg    string      `json:"error,omitempty"`
	ErrorDetail string      `json:"errorDetail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func fail(status int, msg string, detail string) Response {
	return Response{
		Status:      status,
		Success:     false,
		ErrorMsg:    msg,
		ErrorDetail: detail,
	}
}

func ok(data interface{}) Response {
	return Response{
		Status:  200,
		Success: true,
		Data:    data,
	}
}
