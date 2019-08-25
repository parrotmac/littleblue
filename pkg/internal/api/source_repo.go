package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/uuidgen"
)

func (s *apiServer) CreateSourceRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	sourceRepo := &entities.SourceRepository{}

	err := httputils.ReadJsonBodyToEntity(r.Body, sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Overwrite on create
	sourceRepo.RepoUUID = uuidgen.NewUndashed()
	err = s.Storage.CreateSourceRepository(sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (s *apiServer) ListSourceRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	hardcodedUserID := uint(1)

	sourceRepos, err := s.Storage.ListUserSourceRepositories(hardcodedUserID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range sourceRepos {
		sourceRepos[i].AuthenticationCodeSecret = ""
	}

	httputils.RespondWithJSON(w, http.StatusOK, sourceRepos)
}
