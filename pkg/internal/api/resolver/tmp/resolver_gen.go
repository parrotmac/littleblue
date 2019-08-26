package tmp

import (
	"context"

	"github.com/parrotmac/littleblue/pkg/internal/models"
	"github.com/parrotmac/littleblue/pkg/internal/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() server.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() server.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() server.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBuildJob(ctx context.Context, data models.BuildJobCreateInput) (*models.BuildJob, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateBuildJob(ctx context.Context, data models.BuildJobUpdateInput, where models.BuildJobWhereUniqueInput) (*models.BuildJob, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateManyBuildJobs(ctx context.Context, data models.BuildJobUpdateManyMutationInput, where *models.BuildJobWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertBuildJob(ctx context.Context, where models.BuildJobWhereUniqueInput, create models.BuildJobCreateInput, update models.BuildJobUpdateInput) (*models.BuildJob, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteBuildJob(ctx context.Context, where models.BuildJobWhereUniqueInput) (*models.BuildJob, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteManyBuildJobs(ctx context.Context, where *models.BuildJobWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDockerRegistry(ctx context.Context, data models.DockerRegistryCreateInput) (*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateDockerRegistry(ctx context.Context, data models.DockerRegistryUpdateInput, where models.DockerRegistryWhereUniqueInput) (*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateManyDockerRegistries(ctx context.Context, data models.DockerRegistryUpdateManyMutationInput, where *models.DockerRegistryWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertDockerRegistry(ctx context.Context, where models.DockerRegistryWhereUniqueInput, create models.DockerRegistryCreateInput, update models.DockerRegistryUpdateInput) (*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteDockerRegistry(ctx context.Context, where models.DockerRegistryWhereUniqueInput) (*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteManyDockerRegistries(ctx context.Context, where *models.DockerRegistryWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateSourceRepository(ctx context.Context, data models.SourceRepositoryCreateInput) (*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateSourceRepository(ctx context.Context, data models.SourceRepositoryUpdateInput, where models.SourceRepositoryWhereUniqueInput) (*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateManySourceRepositories(ctx context.Context, data models.SourceRepositoryUpdateManyMutationInput, where *models.SourceRepositoryWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertSourceRepository(ctx context.Context, where models.SourceRepositoryWhereUniqueInput, create models.SourceRepositoryCreateInput, update models.SourceRepositoryUpdateInput) (*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteSourceRepository(ctx context.Context, where models.SourceRepositoryWhereUniqueInput) (*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteManySourceRepositories(ctx context.Context, where *models.SourceRepositoryWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, data models.UserCreateInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, data models.UserUpdateInput, where models.UserWhereUniqueInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateManyUsers(ctx context.Context, data models.UserUpdateManyMutationInput, where *models.UserWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertUser(ctx context.Context, where models.UserWhereUniqueInput, create models.UserCreateInput, update models.UserUpdateInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, where models.UserWhereUniqueInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteManyUsers(ctx context.Context, where *models.UserWhereInput) (*models.BatchPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) BuildJob(ctx context.Context, where models.BuildJobWhereUniqueInput) (*models.BuildJob, error) {
	panic("not implemented")
}
func (r *queryResolver) BuildJobs(ctx context.Context, where *models.BuildJobWhereInput, orderBy *models.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*models.BuildJob, error) {
	panic("not implemented")
}
func (r *queryResolver) BuildJobsConnection(ctx context.Context, where *models.BuildJobWhereInput, orderBy *models.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*models.BuildJobConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) DockerRegistry(ctx context.Context, where models.DockerRegistryWhereUniqueInput) (*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *queryResolver) DockerRegistries(ctx context.Context, where *models.DockerRegistryWhereInput, orderBy *models.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*models.DockerRegistry, error) {
	panic("not implemented")
}
func (r *queryResolver) DockerRegistriesConnection(ctx context.Context, where *models.DockerRegistryWhereInput, orderBy *models.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*models.DockerRegistryConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepository(ctx context.Context, where models.SourceRepositoryWhereUniqueInput) (*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepositories(ctx context.Context, where *models.SourceRepositoryWhereInput, orderBy *models.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*models.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepositoriesConnection(ctx context.Context, where *models.SourceRepositoryWhereInput, orderBy *models.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*models.SourceRepositoryConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, where models.UserWhereUniqueInput) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, where *models.UserWhereInput, orderBy *models.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) UsersConnection(ctx context.Context, where *models.UserWhereInput, orderBy *models.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*models.UserConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) Node(ctx context.Context, id string) (models.Node, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) BuildJob(ctx context.Context, where *models.BuildJobSubscriptionWhereInput) (<-chan *models.BuildJobSubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) DockerRegistry(ctx context.Context, where *models.DockerRegistrySubscriptionWhereInput) (<-chan *models.DockerRegistrySubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) SourceRepository(ctx context.Context, where *models.SourceRepositorySubscriptionWhereInput) (<-chan *models.SourceRepositorySubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) User(ctx context.Context, where *models.UserSubscriptionWhereInput) (<-chan *models.UserSubscriptionPayload, error) {
	panic("not implemented")
}
