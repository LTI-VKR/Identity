package errors

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Write(w http.ResponseWriter, r *http.Request, mapped MappedError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(mapped.Status)

	requestID := middleware.GetReqID(r.Context())

	if mapped.Fields != nil {
		_ = json.NewEncoder(w).Encode(ErrorValidationResponse{
			Code:      mapped.Code,
			Fields:    mapped.Fields,
			RequestID: requestID,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(ErrorBasicResponse{
		Code:      mapped.Code,
		Message:   mapped.Message,
		RequestID: requestID,
	})
}
