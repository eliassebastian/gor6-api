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

func (r *queryResolver) Playerquery(ctx context.Context, input model.PlayerSearch) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Searchquery(ctx context.Context, input model.PlayerSearch) ([]*model.PlayerSearchResults, error) {
	log.Println("running searchquery")
	_, l, err := controllers.SearchForPlayer(ctx, r.MC, input.Name, input.Platform)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	fmt.Println("searchquery completed")
	return l, nil
}

func (r *queryResolver) Testquery(ctx context.Context, input model.PlayerSearch) (*model.Player, error) {
	log.Println("running testquery")
	//s, _, err := controllers.SearchForPlayer(ctx, input.Name, input.Platform)
	//if s {
	//	if err != nil {
	//		return nil, gqlerror.Errorf(err.Error())
	//	}
	//	return nil, gqlerror.Errorf("Found Player - Incomplete Function")
	//}

	res, err := controllers.FetchNewPlayer(ctx, r.SM, r.HC, input.Name, input.Platform)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *rankedSeasonResolver) Nextrankmmr(ctx context.Context, obj *model.RankedSeason) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *rankedSeasonResolver) Lastmatchskillstdevchange(ctx context.Context, obj *model.RankedSeason) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *rankedSeasonResolver) Lastmatchskillmeanchange(ctx context.Context, obj *model.RankedSeason) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *rankedSeasonResolver) Previousrankmmr(ctx context.Context, obj *model.RankedSeason) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *rankedSeasonResolver) Pastseasonsabandons(ctx context.Context, obj *model.RankedSeason) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
