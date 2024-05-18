package List

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}
