package funcs

import (
	"math"

	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/model"
)

var PublicTasks = []model.Task{
	{ID: 2, Name: "Задача 2: дифференциальная задача"},
}

const pi = math.Pi

type TaskFactory func(map[string]float64) (model.Task, error)

func GetTask(id int, additionalParams map[string]float64) (model.Task, error) {
	return funcsFactories[id](additionalParams)
}

var funcsFactories map[int]TaskFactory = map[int]TaskFactory{
	2: createProblem2,
}

func createProblem2(params map[string]float64) (model.Task, error) {
	p := model.Task{ID: 2, Name: "Задача 2: дифференциальная задача"}
	p.F = func(x, y float64) float64 {
		return 5 * pi * pi * math.Sin(2*pi*x) * math.Sin(pi*y)
	}
	p.Solution = func(x, y float64) float64 {
		return math.Sin(2*pi*x) * math.Sin(pi*y)
	}
	return p, nil
}
