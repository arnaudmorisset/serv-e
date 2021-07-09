package internal

import "net/http"

const OKResponseBodyMessage string = "OK"

/// write the header status code, header and body message
/// in the response writer as parameter
func formatResponseWriter(
	w http.ResponseWriter,
	statusCode int,
	contentType string,
	bodyMessage []byte,
) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", contentType)
	if len(bodyMessage) > 0 {
		w.Write(bodyMessage)
	}
}

/// format an error response
func returnErrorResponse(w http.ResponseWriter, err error) {
	formatResponseWriter(w, http.StatusInternalServerError, "text/plain", []byte(err.Error()))
}

/// format a response when a record has been created
func returnCreatedRecordResponse(w http.ResponseWriter) {
	formatResponseWriter(w, http.StatusCreated, "text/plain", []byte(OKResponseBodyMessage))
}
