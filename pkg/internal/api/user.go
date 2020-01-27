package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type UserRouter struct {
	StorageService *db.Storage
}

func (router *UserRouter) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newUser := &entities.User{}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(requestBody, newUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = router.StorageService.CreateUser(newUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *UserRouter) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := mux.Vars(r)["user_id"]
	if !ok || userIdStr == "" {
		httputils.RespondWithError(w, http.StatusBadRequest, errors.New("please specify a valid user ID"))
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, errors.New("please specify a valid user ID"))
		return
	}

	user, err := router.StorageService.GetUserByID(uint(userID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, user)
}

func (router *UserRouter) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	updatedUser := &entities.User{}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(requestBody, updatedUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = router.StorageService.UpdateUser(updatedUser)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "updated")
}
