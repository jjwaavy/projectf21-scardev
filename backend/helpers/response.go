package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseType int

const (
	Error ResponseType = iota
	Success
)

type response struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(rt ResponseType, message interface{}, code int, w http.ResponseWriter) {
	var res response

	res.Success = rt == Success

	if rt == Success {
		res.Data = message
	} else if rt == Error {
		res.Error = message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	
	j, _ := json.Marshal(res)

	fmt.Fprintf(w, "%v", string(j))
}
