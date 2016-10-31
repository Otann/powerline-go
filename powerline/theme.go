package powerline

type ColorPair struct {
	Bg string
	Fg string
}

type ColorTriplet struct {
	Bg    string
	Fg    string
	SepFg string
}

type Git struct {
	Clean ColorPair
	Dirty ColorPair
}

type Theme struct {
	ShellBg string
	Host    ColorPair
	Home    ColorPair
	Path    ColorTriplet
	Git
	Lock  ColorPair
	Error ColorPair
}

func SolarizedDark() Theme {
	return Theme{
		ShellBg: "0",
		Host:    ColorPair{Bg: "4", Fg: "7"},
		Home:    ColorPair{Bg: "10", Fg: "0"},
		Path:    ColorTriplet{Bg: "8", Fg: "12", SepFg: "0"},
		Git: Git{
			Clean: ColorPair{Bg: "14", Fg: "0"},
			Dirty: ColorPair{Bg: "2", Fg: "0"},
		},
		Lock:  ColorPair{Bg: "4", Fg: "7"},
		Error: ColorPair{Bg: "1", Fg: "7"},
	}
}
