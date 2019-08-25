package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

func (s *apiServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newUser := &entities.User{}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(requestBody, newUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.Storage.CreateUser(newUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (s *apiServer) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := mux.Vars(r)["user_id"]
	if !ok || userIdStr == "" {
		httputils.RespondWithError(w, http.StatusBadRequest, "please specify a valid user ID")
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, "please specify a valid user ID")
		return
	}

	user, err := s.Storage.GetUserByID(uint(userID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, user)
}

func (s *apiServer) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	updatedUser := &entities.User{}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(requestBody, updatedUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.Storage.UpdateUser(updatedUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "updated")
}
