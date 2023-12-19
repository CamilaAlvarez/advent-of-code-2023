package galaxy

type GalaxyMap [][]string
type Point struct {
	I int
	J int
}
type Galaxy struct {
	Map            GalaxyMap
	GalaxyLocation []Point
}
