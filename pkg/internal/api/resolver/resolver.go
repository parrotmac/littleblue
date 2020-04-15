package resolver

import (
	"context"

	"github.com/parrotmac/littleblue/pkg/internal/client/prisma"
	"github.com/parrotmac/littleblue/pkg/internal/models"
	"github.com/parrotmac/littleblue/pkg/internal/server"
)

/*
Stub is auto-generated. Updates are placed in ./tmp/resolver_gen.go as per ./gqlgen.yml
Copy-paste desired updates into this model
*/

type Resolver struct {
	backend *backend
}

func NewResolver(client *prisma.Client) Resolver {
	return Resolver{
		backend: &backend{
			prisma: client,
		},
	}
}

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

func (r *mutationResolver) CreateBuildJob(ctx context.Context, data prisma.BuildJobCreateInput) (*prisma.BuildJob, error) {
	return r.backend.mutationBackend.CreateBuildJob(ctx, data)
}
func (r *mutationResolver) UpdateBuildJob(ctx context.Context, data prisma.BuildJobUpdateInput, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	return r.backend.mutationBackend.UpdateBuildJob(ctx, data, where)
}
func (r *mutationResolver) UpdateManyBuildJobs(ctx context.Context, data prisma.BuildJobUpdateManyMutationInput, where *prisma.BuildJobWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.UpdateManyBuildJobs(ctx, data, where)
}
func (r *mutationResolver) UpsertBuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput, create prisma.BuildJobCreateInput, update prisma.BuildJobUpdateInput) (*prisma.BuildJob, error) {
	return r.backend.mutationBackend.UpsertBuildJob(ctx, where, create, update)
}
func (r *mutationResolver) DeleteBuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	return r.backend.mutationBackend.DeleteBuildJob(ctx, where)
}
func (r *mutationResolver) DeleteManyBuildJobs(ctx context.Context, where *prisma.BuildJobWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.DeleteManyBuildJobs(ctx, where)
}
func (r *mutationResolver) CreateDockerRegistry(ctx context.Context, data prisma.DockerRegistryCreateInput) (*prisma.DockerRegistry, error) {
	return r.backend.mutationBackend.CreateDockerRegistry(ctx, data)
}
func (r *mutationResolver) UpdateDockerRegistry(ctx context.Context, data prisma.DockerRegistryUpdateInput, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	return r.backend.mutationBackend.UpdateDockerRegistry(ctx, data, where)
}
func (r *mutationResolver) UpdateManyDockerRegistries(ctx context.Context, data prisma.DockerRegistryUpdateManyMutationInput, where *prisma.DockerRegistryWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.UpdateManyDockerRegistries(ctx, data, where)
}
func (r *mutationResolver) UpsertDockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput, create prisma.DockerRegistryCreateInput, update prisma.DockerRegistryUpdateInput) (*prisma.DockerRegistry, error) {
	return r.backend.mutationBackend.UpsertDockerRegistry(ctx, where, create, update)
}
func (r *mutationResolver) DeleteDockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	return r.backend.mutationBackend.DeleteDockerRegistry(ctx, where)
}
func (r *mutationResolver) DeleteManyDockerRegistries(ctx context.Context, where *prisma.DockerRegistryWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.DeleteManyDockerRegistries(ctx, where)
}
func (r *mutationResolver) CreateSourceRepository(ctx context.Context, data prisma.SourceRepositoryCreateInput) (*prisma.SourceRepository, error) {
	return r.backend.mutationBackend.CreateSourceRepository(ctx, data)
}
func (r *mutationResolver) UpdateSourceRepository(ctx context.Context, data prisma.SourceRepositoryUpdateInput, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	return r.backend.mutationBackend.UpdateSourceRepository(ctx, data, where)
}
func (r *mutationResolver) UpdateManySourceRepositories(ctx context.Context, data prisma.SourceRepositoryUpdateManyMutationInput, where *prisma.SourceRepositoryWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.UpdateManySourceRepositories(ctx, data, where)
}
func (r *mutationResolver) UpsertSourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput, create prisma.SourceRepositoryCreateInput, update prisma.SourceRepositoryUpdateInput) (*prisma.SourceRepository, error) {
	return r.backend.mutationBackend.UpsertSourceRepository(ctx, where, create, update)
}
func (r *mutationResolver) DeleteSourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	return r.backend.mutationBackend.DeleteSourceRepository(ctx, where)
}
func (r *mutationResolver) DeleteManySourceRepositories(ctx context.Context, where *prisma.SourceRepositoryWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.DeleteManySourceRepositories(ctx, where)
}
func (r *mutationResolver) CreateUser(ctx context.Context, data prisma.UserCreateInput) (*prisma.User, error) {
	return r.backend.mutationBackend.CreateUser(ctx, data)
}
func (r *mutationResolver) UpdateUser(ctx context.Context, data prisma.UserUpdateInput, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	return r.backend.mutationBackend.UpdateUser(ctx, data, where)
}
func (r *mutationResolver) UpdateManyUsers(ctx context.Context, data prisma.UserUpdateManyMutationInput, where *prisma.UserWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.UpdateManyUsers(ctx, data, where)
}
func (r *mutationResolver) UpsertUser(ctx context.Context, where prisma.UserWhereUniqueInput, create prisma.UserCreateInput, update prisma.UserUpdateInput) (*prisma.User, error) {
	return r.backend.mutationBackend.UpsertUser(ctx, where, create, update)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	return r.backend.mutationBackend.DeleteUser(ctx, where)
}
func (r *mutationResolver) DeleteManyUsers(ctx context.Context, where *prisma.UserWhereInput) (*prisma.BatchPayload, error) {
	return r.backend.mutationBackend.DeleteManyUsers(ctx, where)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) BuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	return r.backend.queryBackend.BuildJob(ctx, where)
}
func (r *queryResolver) BuildJobs(ctx context.Context, where *prisma.BuildJobWhereInput, orderBy *prisma.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.BuildJob, error) {
	return r.backend.queryBackend.BuildJobs(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) BuildJobsConnection(ctx context.Context, where *prisma.BuildJobWhereInput, orderBy *prisma.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.BuildJobConnection, error) {
	return r.backend.queryBackend.BuildJobsConnection(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) DockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	return r.backend.queryBackend.DockerRegistry(ctx, where)
}
func (r *queryResolver) DockerRegistries(ctx context.Context, where *prisma.DockerRegistryWhereInput, orderBy *prisma.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.DockerRegistry, error) {
	return r.backend.queryBackend.DockerRegistries(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) DockerRegistriesConnection(ctx context.Context, where *prisma.DockerRegistryWhereInput, orderBy *prisma.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.DockerRegistryConnection, error) {
	return r.backend.queryBackend.DockerRegistriesConnection(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) SourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	return r.backend.queryBackend.SourceRepository(ctx, where)
}
func (r *queryResolver) SourceRepositories(ctx context.Context, where *prisma.SourceRepositoryWhereInput, orderBy *prisma.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.SourceRepository, error) {
	return r.backend.queryBackend.SourceRepositories(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) SourceRepositoriesConnection(ctx context.Context, where *prisma.SourceRepositoryWhereInput, orderBy *prisma.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.SourceRepositoryConnection, error) {
	return r.backend.queryBackend.SourceRepositoriesConnection(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) User(ctx context.Context, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	return r.backend.queryBackend.User(ctx, where)
}
func (r *queryResolver) Users(ctx context.Context, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.User, error) {
	return r.backend.queryBackend.Users(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) UsersConnection(ctx context.Context, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.UserConnection, error) {
	return r.backend.queryBackend.UsersConnection(ctx, where, orderBy, skip, after, before, first, last)
}
func (r *queryResolver) Node(ctx context.Context, id string) (prisma.Node, error) {
	return r.backend.queryBackend.Node(ctx, id)
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) BuildJob(ctx context.Context, where *prisma.BuildJobSubscriptionWhereInput) (<-chan *prisma.BuildJobSubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) DockerRegistry(ctx context.Context, where *prisma.DockerRegistrySubscriptionWhereInput) (<-chan *prisma.DockerRegistrySubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) SourceRepository(ctx context.Context, where *prisma.SourceRepositorySubscriptionWhereInput) (<-chan *prisma.SourceRepositorySubscriptionPayload, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) User(ctx context.Context, where *prisma.UserSubscriptionWhereInput) (<-chan *prisma.UserSubscriptionPayload, error) {
	panic("not implemented")
}
