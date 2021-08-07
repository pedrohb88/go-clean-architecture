package resolver

import (
	gqlgen "go-clean-architecture/graphql/generated"
	"go-clean-architecture/service"
)

type resolver struct{ service *service.Service }

func New(srvc *service.Service) gqlgen.ResolverRoot {
	return &resolver{service: srvc}
}

func (a *resolver) Mutation() gqlgen.MutationResolver {
	return &mutation{service: a.service}
}

func (a *resolver) Query() gqlgen.QueryResolver {
	return &query{service: a.service}
}
