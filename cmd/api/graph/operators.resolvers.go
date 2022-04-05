package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
)

func (r *operatorSideResolver) Timealivepermatch(ctx context.Context, obj *model.OperatorSide) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *operatorSideResolver) Timedeadpermatch(ctx context.Context, obj *model.OperatorSide) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *operatorSideResolver) Distanceperround(ctx context.Context, obj *model.OperatorSide) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// OperatorSide returns generated.OperatorSideResolver implementation.
func (r *Resolver) OperatorSide() generated.OperatorSideResolver { return &operatorSideResolver{r} }

type operatorSideResolver struct{ *Resolver }
