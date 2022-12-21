package funcs

import "github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/model"

var Schemas = map[int]model.Schema{
	1: {ID: 1, Name: "Метод Якоби (явная схема)", Theta: 0, Tau: 1},
	2: {ID: 2, Name: "Метод Зейделя (релаксационный метод)", Theta: 1, Tau: 1},
	3: {ID: 3, Name: "Своя схема", Theta: 0, Tau: 0},
}

func GetSchema(id int) model.Schema {
	return Schemas[id]
}
