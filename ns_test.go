package ns

import "fmt"

type Solver interface {
	Name() string
	K(F [][]float64, D []float64) [][]float64
}

func Example() {
	fmt.Println("test")
	// Output:
	// test
}
