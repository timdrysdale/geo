package geo

type Point struct {
	X, Y float64
}

type Dim struct {
	W, H float64
}

type Rect struct {
	Corner Point
	Dim    Dim
}

const (
	PXIN = float64(96)                //pixels per inch
	PPIN = float64(72)                //points per inch
	PPMM = float64(PPIN * 1.0 / 25.4) //points per millimetre
	PPPX = float64(PPIN / PXIN)       // points per pixel
)
