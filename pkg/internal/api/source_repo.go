package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/uuidgen"
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

	// Overwrite on create
	sourceRepo.RepoUUID = uuidgen.NewUndashed()
	err = router.StorageService.CreateSourceRepository(sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *SourceRepositoryRouter) ListSourceRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	hardcodedUserID := uint(1)

	sourceRepos, err := router.StorageService.ListUserSourceRepositories(hardcodedUserID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range sourceRepos {
		sourceRepos[i].AuthenticationCodeSecret = ""
	}

	httputils.RespondWithJSON(w, http.StatusOK, sourceRepos)
}
