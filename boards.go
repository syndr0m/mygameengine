package mygameengine

type Boards struct {
	current *Board
	table   map[string]*Board
}

// boards
func (boards *Boards) New(name string) *Board {
	return NewBoard()
}
func (boards *Boards) Register(name string, board *Board) {
	boards.table[name] = board
}
func (boards *Boards) Get(name string) *Board {
	return boards.table[name]
}
func (boards *Boards) SetCurrent(board *Board) {
	if boards.current != nil {
		boards.current.Stop()
	}
	boards.current = board
	boards.current.Start()
}
func (boards *Boards) Current() *Board {
	return boards.current
}

func NewBoards() *Boards {
	boards := new(Boards)
	boards.table = make(map[string]*Board)
	return boards
}
