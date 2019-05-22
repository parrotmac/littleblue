package api

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
)

type SourceRepositoryRouter struct {
	SourceRepositoryService services.SourceRepositoryService
}

func (router *SourceRepositoryRouter) CreateSourceRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	sourceRepo := &storage.SourceRepository{}
	err := httputils.ReadJsonBodyToEntity(r.Body, sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = router.SourceRepositoryService.CreateSourceRepository(sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}
