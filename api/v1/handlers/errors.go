package handlers

import (
	"net/http"

	apimodels "github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/models"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi/operations"
	"github.com/go-openapi/runtime"
)

// DefaultErrorHandler returns an error message when the query parameter is missing.
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	operations.NewGetSearchBadRequest().WithPayload(&apimodels.Error{
		ErrorMessage: "Missing search keyword.",
		ErrorType:    "Client Error",
	}).WriteResponse(w, runtime.JSONProducer())
}
