package tmp

import (
	"context"
	"errors"

	"github.com/parrotmac/littleblue/pkg/models"
	"github.com/parrotmac/littleblue/pkg/prisma"
	"github.com/parrotmac/littleblue/pkg/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var authSecretDummyValue = "********************"

type Resolver struct {
	Prisma *prisma.Client
}

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
	return r.Prisma.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
}
func (r *mutationResolver) EnableSourceRepository(ctx context.Context, repoID string) (*prisma.SourceRepository, error) {
	return r.Prisma.SourceRepository(prisma.SourceRepositoryWhereUniqueInput{
		ID: &repoID,
	}).Exec(ctx)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserByID(ctx context.Context, id string) (*prisma.User, error) {
	return r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)
}
func (r *queryResolver) EnabledRepositories(ctx context.Context) ([]*prisma.SourceRepository, error) {
	panic("not implemented")
}
func (r *queryResolver) SourceRepositoryByOwnerAndName(ctx context.Context, ownerID string, name string) (*prisma.SourceRepository, error) {
	repos, err := r.Prisma.SourceRepositories(&prisma.SourceRepositoriesParams{
		Where: &prisma.SourceRepositoryWhereInput{
			Name: &name,
			Owner: &prisma.UserWhereInput{
				ID: &ownerID,
			},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	if len(repos) == 0 || len(repos) > 1 {
		return nil, errors.New("no match found")
	}
	repo := repos[0]
	repo.AuthSecret = &authSecretDummyValue
	return &repo, nil
}
func (r *queryResolver) SourceRepositoryByID(ctx context.Context, repoID string) (*prisma.SourceRepository, error) {
	repo, err := r.Prisma.SourceRepository(prisma.SourceRepositoryWhereUniqueInput{
		ID: &repoID,
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	repo.AuthSecret = &authSecretDummyValue
	return repo, nil
}
func (r *queryResolver) UserSourceRepositories(ctx context.Context, userID string) (repos []*prisma.SourceRepository, err error) {
	sourceRepositories, err := r.Prisma.SourceRepositories(&prisma.SourceRepositoriesParams{
		Where: &prisma.SourceRepositoryWhereInput{
			Owner: &prisma.UserWhereInput{
				ID: &userID,
			},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, sourceRepo := range sourceRepositories {
		sourceRepo.AuthSecret = &authSecretDummyValue
		repos = append(repos, &sourceRepo)
	}
	return repos, err
}

type sourceRepositoryResolver struct{ *Resolver }

func (r *sourceRepositoryResolver) Owner(ctx context.Context, obj *prisma.SourceRepository) (*prisma.User, error) {
	return r.Prisma.SourceRepository(prisma.SourceRepositoryWhereUniqueInput{
		ID: &obj.ID,
	}).Owner().Exec(ctx)
}

func (r *sourceRepositoryResolver) SourceProvider(ctx context.Context, obj *prisma.SourceRepository) (*models.SourceProvider, error) {
	repo, err := r.Prisma.SourceRepository(prisma.SourceRepositoryWhereUniqueInput{
		ID: &obj.ID,
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	provider := models.SourceProvider(*repo.SourceProvider)
	return &provider, nil
}

func (r *sourceRepositoryResolver) CloneStrategy(ctx context.Context, obj *prisma.SourceRepository) (models.CloneStrategy, error) {
	repo, err := r.Prisma.SourceRepository(prisma.SourceRepositoryWhereUniqueInput{
		ID: &obj.ID,
	}).Exec(ctx)
	if err != nil {
		return models.CloneStrategy(""), err
	}
	return models.CloneStrategy(repo.CloneStrategy), nil
}
