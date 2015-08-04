package mygameengine

type Message struct {
	name  int
	value int
}

const (
	MESSAGE_EXIT = iota
	MESSAGE_REPAINT
	MESSAGE_KEY_DOWN
	MESSAGE_KEY_UP
)
