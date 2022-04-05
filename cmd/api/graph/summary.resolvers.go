package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
)

func (r *summarySeasonResolver) Timealivepermatch(ctx context.Context, obj *model.SummarySeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *summarySeasonResolver) Timedeadpermatch(ctx context.Context, obj *model.SummarySeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *summarySeasonResolver) Distanceperround(ctx context.Context, obj *model.SummarySeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// SummarySeason returns generated.SummarySeasonResolver implementation.
func (r *Resolver) SummarySeason() generated.SummarySeasonResolver { return &summarySeasonResolver{r} }

type summarySeasonResolver struct{ *Resolver }
