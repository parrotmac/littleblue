package api

import (
	"github.com/gin-gonic/gin"

	"github.com/parrotmac/littleblue/pkg/internal/builder"
	"github.com/parrotmac/littleblue/pkg/internal/db"
)

type apiServer struct {
	RouteGroup *gin.RouterGroup
	Storage    *db.Storage
	Builder    *builder.Builder
}

func NewServer(routeGroup *gin.RouterGroup, storage *db.Storage, builder *builder.Builder) *apiServer {
	s := &apiServer{
		RouteGroup: routeGroup,
		Storage:    storage,
		Builder:    builder,
	}

	s.initRoutes()
	return s
}

func (s *apiServer) initRoutes() {
	s.RouteGroup.POST("/webhook/{repo_uuid}/", gin.WrapF(s.WebhookJobHandler))
}
