package httputils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithStatus(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"status": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		// This is (at least theoretically) dangerous as respondWithError calls this function
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		logrus.Errorln(err)
	}
}

func ReadJsonBodyToEntity(body io.ReadCloser, dst interface{}) error {
	bodyData, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyData, dst)
}
