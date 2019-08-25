package tmp

import (
	"context"

	"github.com/parrotmac/littleblue/pkg/models"
	"github.com/parrotmac/littleblue/pkg/prisma"
	"github.com/parrotmac/littleblue/pkg/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() server.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() server.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) SourceRepository() server.SourceRepositoryResolver {
	return &sourceRepositoryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) EnableSourceRepository(ctx context.Context, repoID string) (*prisma.SourceRepository, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserByID(ctx context.Context, id string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) EnabledRepositories(ctx context.Context) ([]*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepositoryByOwnerAndName(ctx context.Context, ownerID string, name string) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepositoryByID(ctx context.Context, repoID string) (*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) UserSourceRepositories(ctx context.Context, userID string) ([]*prisma.SourceRepository, error) {
	panic("not implemented")
}

type sourceRepositoryResolver struct{ *Resolver }

func (r *sourceRepositoryResolver) Owner(ctx context.Context, obj *prisma.SourceRepository) (*prisma.User, error) {
	panic("not implemented")
}
func (r *sourceRepositoryResolver) SourceProvider(ctx context.Context, obj *prisma.SourceRepository) (*models.SourceProvider, error) {
	panic("not implemented")
}
func (r *sourceRepositoryResolver) CloneStrategy(ctx context.Context, obj *prisma.SourceRepository) (models.CloneStrategy, error) {
	panic("not implemented")
}
