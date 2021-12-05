package resolver

import (
	"context"
	"go_gorm_graphql/db"
	"go_gorm_graphql/model"
)

type Resolver struct{}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func New() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	db := db.GetDB()

	account := model.Account{
		ID:   "1",
		Name: input.Name,
	}

	if err := db.Create(&account).Error; err != nil {
		return &model.Account{}, err
	}

	return &account, nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	var accounts []*model.Account
	db := db.GetDB()
	db.Find(&accounts)
	return accounts, nil
}
