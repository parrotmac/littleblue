package resolver

import (
	"context"

	"github.com/parrotmac/littleblue/pkg/internal/client/prisma"
)

type queryBackend struct {
	*backend
}

func (b *queryBackend) BuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	return b.prisma.BuildJob(where).Exec(ctx)
}
func (b *queryBackend) BuildJobs(ctx context.Context, where *prisma.BuildJobWhereInput, orderBy *prisma.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.BuildJob, error) {
	panic("not implemented")
}
func (b *queryBackend) BuildJobsConnection(ctx context.Context, where *prisma.BuildJobWhereInput, orderBy *prisma.BuildJobOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.BuildJobConnection, error) {
	panic("not implemented")
}
func (b *queryBackend) DockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *queryBackend) DockerRegistries(ctx context.Context, where *prisma.DockerRegistryWhereInput, orderBy *prisma.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *queryBackend) DockerRegistriesConnection(ctx context.Context, where *prisma.DockerRegistryWhereInput, orderBy *prisma.DockerRegistryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.DockerRegistryConnection, error) {
	panic("not implemented")
}
func (b *queryBackend) SourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *queryBackend) SourceRepositories(ctx context.Context, where *prisma.SourceRepositoryWhereInput, orderBy *prisma.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *queryBackend) SourceRepositoriesConnection(ctx context.Context, where *prisma.SourceRepositoryWhereInput, orderBy *prisma.SourceRepositoryOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.SourceRepositoryConnection, error) {
	panic("not implemented")
}
func (b *queryBackend) User(ctx context.Context, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	panic("not implemented")
}
func (b *queryBackend) Users(ctx context.Context, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.User, error) {
	panic("not implemented")
}
func (b *queryBackend) UsersConnection(ctx context.Context, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) (*prisma.UserConnection, error) {
	panic("not implemented")
}
func (b *queryBackend) Node(ctx context.Context, id string) (prisma.Node, error) {
	panic("not implemented")
}
