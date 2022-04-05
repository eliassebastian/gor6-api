package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
)

func (r *mapResolver) Timealivepermatch(ctx context.Context, obj *model.Map) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mapResolver) Timedeadpermatch(ctx context.Context, obj *model.Map) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mapResolver) Distanceperround(ctx context.Context, obj *model.Map) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Map returns generated.MapResolver implementation.
func (r *Resolver) Map() generated.MapResolver { return &mapResolver{r} }

type mapResolver struct{ *Resolver }
