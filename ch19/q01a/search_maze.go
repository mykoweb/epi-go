package ch19q01a

type Coord struct{ x, y int }

// A maze is represented by a graph of Squares. The bottom left corner is
// coord 1,1. The square to its right is 1,2.
type Square struct {
	location Coord
	adj      []*Square
	visited  bool
	finish   bool
}

// SearchMaze searches a maze represented by the Square struct. Each Square
// can be viewed as a graph vertex. An edge between 2 adjacent squares means
// there is a path between them.
//
// SearchMaze accepts a pointer to a Square and this represents the start of the
// Maze. The finish point, if it exists, is represented by a Square whose bool
// attribute, finish, is set to true.
//
// SearchMaze employs depth-first search and, thus, is not guaranteed to find
// the shortest path to the finish point.
func SearchMaze(s *Square) bool {
	if (*s).finish {
		return true
	}
	(*s).visited = true
	foundFinish := false

	for _, adjSquare := range (*s).adj {
		if !(*adjSquare).visited {
			foundFinish = foundFinish || SearchMaze(adjSquare)
		}
	}

	return foundFinish
}
