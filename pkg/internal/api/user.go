package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type UserRouter struct {
	UserService services.UserService
}

func (router *UserRouter) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newUser := &storage.User{}

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

	err = router.UserService.CreateUser(newUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *UserRouter) GetUserHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := router.UserService.GetUserByID(uint(userID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, user)
}

func (router *UserRouter) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	updatedUser := &storage.User{}

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

	err = router.UserService.UpdateUser(updatedUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "updated")
}
