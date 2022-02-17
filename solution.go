package square

import "math"

type myInt int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * sideLen * sideLen
	case 3:
		return math.Sqrt(3) * sideLen * sideLen / 4
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
