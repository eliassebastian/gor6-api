package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eliassebastian/gor6-api/cmd/api/graph/generated"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
)

func (r *rankedSeasonResolver) Maxmmr(ctx context.Context, obj *model.RankedSeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rankedSeasonResolver) Skillmean(ctx context.Context, obj *model.RankedSeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rankedSeasonResolver) Skillstdev(ctx context.Context, obj *model.RankedSeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rankedSeasonResolver) Lastmatchmmrchange(ctx context.Context, obj *model.RankedSeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rankedSeasonResolver) Mmr(ctx context.Context, obj *model.RankedSeason) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// RankedSeason returns generated.RankedSeasonResolver implementation.
func (r *Resolver) RankedSeason() generated.RankedSeasonResolver { return &rankedSeasonResolver{r} }

type rankedSeasonResolver struct{ *Resolver }
