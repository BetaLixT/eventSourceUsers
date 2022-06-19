package repo

type Event struct {
	Id            string
	Stream        string
	StreamId      string
	StreamVersion int
	Data          interface{}
}
