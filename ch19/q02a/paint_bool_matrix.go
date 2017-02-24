package ch19q02a

type Square struct {
	X, Y    int
	Val     bool
	Adj     [4]*Square // {left, right, up, down} adj[0] is left
	Visited bool
}

// PaintBoolMatrix takes a pointer to a square within an nxm boolean matrix and
// flips the boolean value of the squares associated with the given square.
func PaintBoolMatrix(start *Square) {
	queue := []*Square{}
	(*start).Visited = true
	color := (*start).Val
	(*start).Val = !color
	queue = append(queue, start)

	for len(queue) != 0 {
		currSquare := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		for _, sq := range currSquare.Adj {
			if sq != nil && !(*sq).Visited && (*sq).Val == color {
				(*sq).Val = !color
				(*sq).Visited = true
				queue = append(queue, sq)
			}
		}
	}
}
