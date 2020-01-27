package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type SourceProviderRouter struct {
	StorageService *db.Storage
}

func (router *SourceProviderRouter) CreateSourceProviderHandler(w http.ResponseWriter, r *http.Request) {

	// FIXME: Lookup user from session
	hardCodedOwnerUserId := uint(1)

	sourceProvider := &entities.SourceProvider{
		OwnerID: hardCodedOwnerUserId,
	}

	err := httputils.ReadJsonBodyToEntity(r.Body, sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = router.StorageService.CreateSourceProvider(sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *SourceProviderRouter) ListSourceProvidersHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Lookup user from session
	hardCodedOwnerUserId := uint(1)

	sourceProviders, err := router.StorageService.ListUserSourceProviders(hardCodedOwnerUserId)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Do this automatically
	for i := range sourceProviders {
		sourceProviders[i].AuthorizationToken = ""
	}

	httputils.RespondWithJSON(w, http.StatusOK, sourceProviders)
}
