package color

type Color struct{}

type RGB struct {
	Red   uint8
	Green uint8
	Bluee uint8
}

type RGBA struct {
	RGB
	Alpha uint8
}

type CMYK struct {
	Cyan    uint8
	Magenta uint8
	Yellow  uint8
	Black   uint8
}

type HSL struct {
	Hue        uint16 // 0 - 360
	Saturation uint16 // 0% - 100%
	Luminance  uint16 // 0% - 100%
}

func (rgb RGB) Scan(input string) {
	// mask #RRGGBB
}

func (rgba RGBA) Scan(input string) {
	// mask #RRGGBBAA
}

func (hsl HSL) Scan(input string) {
	// mask ((H || HH || HHH), (S || SS || 100)%, ( L || LL || 100)%)
	// (240, 100%, 50%) == #0000FF
}
