package ns

type Solver interface {
	Name() string
	K(F [][]float64, D []float64) [][]float64
}
