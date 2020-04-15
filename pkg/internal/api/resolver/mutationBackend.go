package resolver

import (
	"context"

	"github.com/parrotmac/littleblue/pkg/internal/client/prisma"
	"github.com/parrotmac/littleblue/pkg/internal/models"
)

type mutationBackend struct{ *Resolver }

func (b *mutationBackend) CreateBuildJob(ctx context.Context, data prisma.BuildJobCreateInput) (*prisma.BuildJob, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateBuildJob(ctx context.Context, data prisma.BuildJobUpdateInput, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateManyBuildJobs(ctx context.Context, data prisma.BuildJobUpdateManyMutationInput, where *prisma.BuildJobWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpsertBuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput, create prisma.BuildJobCreateInput, update prisma.BuildJobUpdateInput) (*prisma.BuildJob, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteBuildJob(ctx context.Context, where prisma.BuildJobWhereUniqueInput) (*prisma.BuildJob, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteManyBuildJobs(ctx context.Context, where *prisma.BuildJobWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) CreateDockerRegistry(ctx context.Context, data prisma.DockerRegistryCreateInput) (*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateDockerRegistry(ctx context.Context, data prisma.DockerRegistryUpdateInput, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateManyDockerRegistries(ctx context.Context, data prisma.DockerRegistryUpdateManyMutationInput, where *prisma.DockerRegistryWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpsertDockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput, create prisma.DockerRegistryCreateInput, update prisma.DockerRegistryUpdateInput) (*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteDockerRegistry(ctx context.Context, where prisma.DockerRegistryWhereUniqueInput) (*prisma.DockerRegistry, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteManyDockerRegistries(ctx context.Context, where *prisma.DockerRegistryWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) CreateSourceRepository(ctx context.Context, data prisma.SourceRepositoryCreateInput) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateSourceRepository(ctx context.Context, data prisma.SourceRepositoryUpdateInput, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateManySourceRepositories(ctx context.Context, data prisma.SourceRepositoryUpdateManyMutationInput, where *prisma.SourceRepositoryWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpsertSourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput, create prisma.SourceRepositoryCreateInput, update prisma.SourceRepositoryUpdateInput) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteSourceRepository(ctx context.Context, where prisma.SourceRepositoryWhereUniqueInput) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteManySourceRepositories(ctx context.Context, where *prisma.SourceRepositoryWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) CreateUser(ctx context.Context, data prisma.UserCreateInput) (*prisma.User, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateUser(ctx context.Context, data prisma.UserUpdateInput, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpdateManyUsers(ctx context.Context, data prisma.UserUpdateManyMutationInput, where *prisma.UserWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
func (b *mutationBackend) UpsertUser(ctx context.Context, where prisma.UserWhereUniqueInput, create prisma.UserCreateInput, update prisma.UserUpdateInput) (*prisma.User, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteUser(ctx context.Context, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	panic("not implemented")
}
func (b *mutationBackend) DeleteManyUsers(ctx context.Context, where *prisma.UserWhereInput) (*prisma.BatchPayload, error) {
	panic("not implemented")
}
