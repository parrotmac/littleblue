package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

func (s *apiServer) CreateSourceProviderHandler(w http.ResponseWriter, r *http.Request) {

	// FIXME: Lookup user from session
	hardCodedOwnerUserId := uint(1)

	sourceProvider := &entities.SourceProvider{
		OwnerID: hardCodedOwnerUserId,
	}

	err := httputils.ReadJsonBodyToEntity(r.Body, sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.Storage.CreateSourceProvider(sourceProvider)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (s *apiServer) ListSourceProvidersHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Lookup user from session
	hardCodedOwnerUserId := uint(1)

	sourceProviders, err := s.Storage.ListUserSourceProviders(hardCodedOwnerUserId)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Do this automatically
	for i := range sourceProviders {
		sourceProviders[i].AuthorizationToken = ""
	}

	httputils.RespondWithJSON(w, http.StatusOK, sourceProviders)
}
