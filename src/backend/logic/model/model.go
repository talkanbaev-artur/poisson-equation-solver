package model

type Numericals struct {
	MethodID          int     `json:"id"`
	MethodTitle       string  `json:"title"`
	MethodDescription string  `json:"description"`
	MethodIcon        *string `json:"icon"`
}

type SolutionParameters struct {
	GridSize             int                `json:"n"`
	TaskID               int                `json:"task"`
	AdditionalParameters map[string]float64 `json:"special"`
	Theta                float64            `json:"theta"`
	Tau                  float64            `json:"tau"`
}

type SolutionData struct {
	Error          float64     `json:"err"`
	Sigma          float64     `json:"sigma"`
	IterationCount int         `json:"k"`
	Solution       [][]float64 `json:"data"`
}

type Schema struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Theta float64 `json:"theta"`
	Tau   float64 `json:"tau"`
}

type RF func(x float64) float64

type D2RF func(x, t float64) float64

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	F        D2RF   `json:"-"`
	Solution D2RF   `json:"-"`
}
