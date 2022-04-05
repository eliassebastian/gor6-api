package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
)

func (r *primaryWeaponsResolver) Weapontypes(ctx context.Context, obj *model.PrimaryWeapons) (*model.WeaponTypes, error) {
	panic(fmt.Errorf("not implemented"))
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

// SecondaryWeapons returns generated.SecondaryWeaponsResolver implementation.
func (r *Resolver) SecondaryWeapons() generated.SecondaryWeaponsResolver {
	return &secondaryWeaponsResolver{r}
}

// Weapons returns generated.WeaponsResolver implementation.
func (r *Resolver) Weapons() generated.WeaponsResolver { return &weaponsResolver{r} }

type primaryWeaponsResolver struct{ *Resolver }
type secondaryWeaponsResolver struct{ *Resolver }
type weaponsResolver struct{ *Resolver }
