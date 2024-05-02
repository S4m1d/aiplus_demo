package api

import (
	"aiplus_demo/src/bl"
	"aiplus_demo/src/da"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

const reqHandleErrTmplt = "failed to handle request %s %s"

func HandlePostEmployeesData(responseWriter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var empl da.Employee
	err := decoder.Decode(&empl)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write(nil)
		return
	}

	ctx, cancel := context.WithTimeout(request.Context(), 10*time.Second)
	defer cancel()

	isSuccess, wErr := bl.EmplWriter.Write(ctx, empl)
	if wErr != nil {
		log.Error().Err(err).Msgf(reqHandleErrTmplt, request.Method, request.RequestURI)
		switch wErr.ErrType {
		case bl.BlValidationError:
			responseWriter.WriteHeader(http.StatusBadRequest)
		default:
			responseWriter.WriteHeader(http.StatusInternalServerError)
		}
		responseWriter.Write(nil)
	}

	responseWriter.WriteHeader(http.StatusOK)
	respBody := PostEmployeesDataRs{
		Success: isSuccess,
	}
	mrshBody, err := json.Marshal(respBody)
	if err != nil {
		log.Error().Err(err).Msgf(reqHandleErrTmplt, request.Method, request.RequestURI)
	}
	responseWriter.Write(mrshBody)
}
