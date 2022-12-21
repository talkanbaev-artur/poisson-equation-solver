package algorithm

import (
	"math"

	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic/model"
)

func Solve(task model.Task, params model.SolutionParameters) *model.SolutionData {
	//init solution
	sol := model.SolutionData{}
	//define borders
	sol.Sigma = 1 / math.Pow(float64(params.GridSize-1), 3)
	n, h, th, t := params.GridSize*1.0, 1.0/float64(params.GridSize-1), params.Theta, params.Tau
	//loop
	err := math.Inf(1)
	u := [][][]float64{}
	f := [][]float64{}
	o := [][]float64{}
	{
		s := [][]float64{}
		for j := 0; j < n; j++ {
			s = append(s, make([]float64, n))
			f = append(f, make([]float64, n))
			o = append(o, make([]float64, n))
			for i := 0; i < n; i++ {
				s[j][i] = 0
				f[j][i] = task.F(float64(i)*h, float64(j)*h)
				o[j][i] = task.Solution(float64(i)*h, float64(j)*h)
			}
		}
		u = append(u, s)
	}

	tth, th4, t4, th24, t1 := (t-th)/4, th/4, t/4, t*h*h/4, 1-t
	k := 1
	for k = 1; k < 1000; k++ {
		//create solution grid
		s := [][]float64{}
		for i := 0; i < n; i++ {
			s = append(s, make([]float64, n))
		}
		maxErr := 0.0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == 0 || j == 0 || i == n-1 || j == n-1 {
					s[j][i] = 0
					continue
				}
				s[j][i] = th4*(s[j][i-1]+s[j-1][i]) +
					tth*(u[k-1][j][i-1]+u[k-1][j-1][i]) +
					t4*(u[k-1][j][i+1]+u[k-1][j+1][i]) +
					t1*u[k-1][j][i] +
					th24*f[j][i]
				if err := math.Abs(o[j][i] - s[j][i]); err > maxErr {
					maxErr = err
				}
			}
		}
		u = append(u, s)
		err = maxErr
		cheb := 0.0
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				if err := math.Abs(u[k][j][i] - u[k-1][j][i]); err > cheb {
					cheb = err
				}
			}
		}
		if cheb <= sol.Sigma {
			break
		}
	}
	sol.Orginal = o
	sol.Error = err
	sol.IterationCount = k
	sol.Solution = u
	return &sol
}
