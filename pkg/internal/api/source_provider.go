package api

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
)

type SourceProviderRouter struct {
	SourceProviderService services.SourceProviderService
}

func (router *SourceProviderRouter) CreateSourceProviderHandler(w http.ResponseWriter, r *http.Request) {
	sourceProvider := &storage.SourceProvider{}
	err := httputils.ReadJsonBodyToEntity(r.Body, sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = router.SourceProviderService.CreateSourceProvider(sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}
