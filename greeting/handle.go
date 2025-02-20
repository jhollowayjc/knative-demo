package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Message string
}

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	req := request{
		Name: "friend",
	}

	// TODO: handle err
	json.NewDecoder(r.Body).Decode(&req)

	rsp := response{
		Message: fmt.Sprintf("Hello %s, nice to meet you", req.Name),
	}

	// TODO: handle err
	json.NewEncoder(w).Encode(rsp)
}
