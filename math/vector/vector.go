package vector

type Vector struct {
	Dim  int32
	Data []float64
}

type Matrix struct {
	DimX int32
	DimY int32
	Data []Vector
}
