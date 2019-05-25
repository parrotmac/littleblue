package api

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type SourceRepositoryRouter struct {
	StorageService *db.Storage
}

func (router *SourceRepositoryRouter) CreateSourceRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	sourceRepo := &entities.SourceRepository{}

	err := httputils.ReadJsonBodyToEntity(r.Body, sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Creating Source Repo: %+v", *sourceRepo)

	err = router.StorageService.CreateSourceRepository(sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}
