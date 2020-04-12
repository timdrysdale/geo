package geo

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Dim struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type Rect struct {
	Corner Point `json:"corner"`
	Dim    Dim   `json:"dim"`
}

type DynamicDim struct {
	Dim             Dim  `json:"dim"`
	WidthIsDynamic  bool `json:"widthIsDynamic"`
	HeightIsDynamic bool `json:"heightIsDynamic"`
}

const (
	AnchorsLayer    = "anchors"
	AnchorReference = "ref-anchor"
	ChromeLayer     = "chrome"
	TextFieldsLayer = "textfields"
	PagesLayer      = "pages"
	ImagesLayer     = "images"
	Translate       = "translate"
	PXIN            = float64(96)                //pixels per inch
	PPIN            = float64(72)                //points per inch
	PPMM            = float64(PPIN * 1.0 / 25.4) //points per millimetre
	PPPX            = float64(PPIN / PXIN)       // points per pixel

)
