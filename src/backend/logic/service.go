package logic

import (
	"context"

	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/funcs"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/model"
)

type APIService interface {
	GetNumericals(ctx context.Context) []model.Numericals
	Solve(ctx context.Context, params model.SolutionParameters) (model.SolutionData, error)
	GetTasks(ctx context.Context) []model.Task
	GetSchemas(ctx context.Context) []model.Schema
}

func NewAPIService() APIService {
	return service{}
}

type service struct {
}

func (s service) GetNumericals(ctx context.Context) []model.Numericals {
	nums := []model.Numericals{
		{MethodID: 1, MethodTitle: "Numerical methods for 2nd-order ODE Boundary problems"},
	}

	return nums
}

func (s service) Solve(ctx context.Context, params model.SolutionParameters) (model.SolutionData, error) {
	d := model.SolutionData{Solution: make([][]float64, 0)}
	d.Solution = append(d.Solution, []float64{1, 2, 3}, []float64{1, 2, 3}, []float64{1, 2, 3})
	return d, nil
}

func (s service) GetSchemas(ctx context.Context) []model.Schema {
	schemas := make([]model.Schema, len(funcs.Schemas))
	for i, s := range funcs.Schemas {
		schemas[i-1] = s
	}
	return schemas
}

func (s service) GetTasks(ctx context.Context) []model.Task {
	return funcs.PublicTasks
}
