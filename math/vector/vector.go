package vector

import "fmt"

type Vector struct {
	Dim  int
	Data []float64
}

func New(dim int, data []float64) *Vector {
	return &Vector{
		Dim:  dim,
		Data: data,
	}
}

func (v *Vector) Scalar(k float64) {
	for index, _ := range v.Data {
		v.Data[index] *= k
	}
}

func (v *Vector) Add(vector *Vector) error {
	if v.Dim != vector.Dim {
		return fmt.Errorf("unmatched dimensions: %v and %v", v.Dim, vector.Dim)
	}

	for index, value := range vector.Data {
		v.Data[index] += value
	}

	return nil
}

func (v *Vector) Sub(vector *Vector) error {
	if v.Dim != vector.Dim {
		return fmt.Errorf("unmatched dimensions: %v and %v", v.Dim, vector.Dim)
	}

	for index, value := range vector.Data {
		v.Data[index] -= value
	}

	return nil
}

func (v *Vector) Mul(vector *Vector) error {
	if v.Dim != vector.Dim {
		return fmt.Errorf("unmatched dimensions: %v and %v", v.Dim, vector.Dim)
	}

	for index, value := range vector.Data {
		v.Data[index] *= value
	}

	return nil
}
