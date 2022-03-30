package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/eliassebastian/gor6-api/cmd/api/controllers"
	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *primaryWeaponsResolver) Weapontypes(ctx context.Context, obj *model.PrimaryWeapons) (*model.WeaponTypes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Playerquery(ctx context.Context, input model.PlayerSearch) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Searchquery(ctx context.Context, input model.PlayerSearch) ([]*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Testquery(ctx context.Context, input model.PlayerSearch) (*model.Player, error) {
	log.Println("running testquery")
	s, _, err := controllers.SearchForPlayer(ctx, input.Name, input.Platform)
	if s {
		if err != nil {
			return nil, gqlerror.Errorf(err.Error())
		}
		return nil, gqlerror.Errorf("Found Player - Incomplete Function")
	}

	res, err := controllers.FetchNewPlayer(ctx, r.MC, r.SM, r.HC, input.Name, input.Platform)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return res, nil
}

func (r *secondaryWeaponsResolver) Weapontypes(ctx context.Context, obj *model.SecondaryWeapons) (*model.WeaponTypes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *weaponsResolver) Roundswithmultikill(ctx context.Context, obj *model.Weapons) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// PrimaryWeapons returns generated.PrimaryWeaponsResolver implementation.
func (r *Resolver) PrimaryWeapons() generated.PrimaryWeaponsResolver {
	return &primaryWeaponsResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SecondaryWeapons returns generated.SecondaryWeaponsResolver implementation.
func (r *Resolver) SecondaryWeapons() generated.SecondaryWeaponsResolver {
	return &secondaryWeaponsResolver{r}
}

// Weapons returns generated.WeaponsResolver implementation.
func (r *Resolver) Weapons() generated.WeaponsResolver { return &weaponsResolver{r} }

type primaryWeaponsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type secondaryWeaponsResolver struct{ *Resolver }
type weaponsResolver struct{ *Resolver }
