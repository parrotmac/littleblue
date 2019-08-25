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
	s.RouteGroup.POST("/users/", gin.WrapF(s.CreateUserHandler))
	s.RouteGroup.GET("/users/{user_id}/", gin.WrapF(s.GetUserHandler))
	s.RouteGroup.PATCH("/users/{user_id}/", gin.WrapF(s.UpdateUserHandler))
	s.RouteGroup.POST("/registries/", gin.WrapF(s.CreateDockerRegistryHandler))
	s.RouteGroup.POST("/source-providers/", gin.WrapF(s.CreateSourceProviderHandler))
	s.RouteGroup.GET("/source-providers/", gin.WrapF(s.ListSourceProvidersHandler))
	s.RouteGroup.POST("/jobs/", gin.WrapF(s.CreateBuildJobHandler))
	s.RouteGroup.POST("/webhook/{repo_uuid}/", gin.WrapF(s.WebhookJobHandler))
	s.RouteGroup.POST("/repos/", gin.WrapF(s.CreateSourceRepositoryHandler))
	s.RouteGroup.GET("/repos/", gin.WrapF(s.ListSourceRepositoriesHandler))
	s.RouteGroup.POST("/repos/{repo_id}/configs/", gin.WrapF(s.CreateBuildConfigHandler))
	s.RouteGroup.GET("/repos/{repo_id}/configs/", gin.WrapF(s.ListRepoBuildConfigurationsHandler))
	s.RouteGroup.GET("/repos/{repo_uuid}/jobs/", gin.WrapF(s.ListRepoBuildJobsHandler))
}
