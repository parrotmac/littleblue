package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
)

type SourceRepositoryRouter struct {
	SourceRepositoryService services.SourceRepositoryService
}

func (router *SourceRepositoryRouter) CreateSourceRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	httputils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
