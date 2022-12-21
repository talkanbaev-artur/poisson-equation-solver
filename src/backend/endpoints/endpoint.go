package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/model"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/util"
)

type Endpoints struct {
	GetNumericals      endpoint.Endpoint
	SolveEndpoint      endpoint.Endpoint
	GetTasksEndpoint   endpoint.Endpoint
	GetSchemasEndpoint endpoint.Endpoint
}

func CreateEndpoints(s logic.APIService, lg log.Logger) Endpoints {
	es := Endpoints{}
	es.GetNumericals = MakeGetNumericalsEndpoint(s, lg)
	es.GetSchemasEndpoint = makeGetSchemasEndpoint(s, lg)
	es.GetTasksEndpoint = makeGetTasksEndpoint(s, lg)
	es.SolveEndpoint = makeSolveEndpoint(s, lg)
	return es
}

func MakeGetNumericalsEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetNumericals(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get numers")(e)
	return e
}

func makeSolveEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		p := request.(model.SolutionParameters)
		return s.Solve(ctx, p)
	}
	e = util.LoggingMiddleware(lg, "solve")(e)
	return e
}

func makeGetTasksEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetTasks(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get tasks")(e)
	return e
}

func makeGetSchemasEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetSchemas(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get schemas")(e)
	return e
}
