package pipemap

type Pipe int

const (
	S Pipe = iota
	I
	H
	L
	J
	Seven
	F
	Dot
)

type PipeMap struct {
	Map              [][]Pipe
	StartLocationRow int
	StartLocationCol int
}
