package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

func (s *apiServer) CreateDockerRegistryHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME Lookup user
	hardcodedUserID := uint(1)

	registry := &entities.DockerRegistry{}
	err := httputils.ReadJsonBodyToEntity(r.Body, registry)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	registry.OwnerID = hardcodedUserID // FIXME

	err = s.Storage.CreateDockerRegistry(registry)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	registry.Password = ""

	httputils.RespondWithJSON(w, http.StatusCreated, registry)
}
