package matrix

import "github.com/medtune/go-utils/math/vector"

type Matrix struct {
	DimX int32
	DimY int32
	Data []vector.Vector
}
