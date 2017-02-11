package ch19q01a

import "testing"

func TestSearchMaze(t *testing.T) {
	badMaze := Square{location: Coord{1, 1}}
	sq1_2 := Square{location: Coord{1, 2}}
	sq2_2 := Square{location: Coord{2, 2}}
	sq2_3 := Square{location: Coord{2, 3}}
	sq3_2 := Square{location: Coord{3, 2}}

	badMaze.adj = append(badMaze.adj, &sq1_2)
	sq1_2.adj = append(sq1_2.adj, &badMaze)
	sq1_2.adj = append(sq1_2.adj, &sq2_2)
	sq2_2.adj = append(sq2_2.adj, &sq1_2)
	sq2_2.adj = append(sq2_2.adj, &sq2_3)
	sq2_2.adj = append(sq2_2.adj, &sq3_2)
	sq2_3.adj = append(sq2_3.adj, &sq2_2)
	sq3_2.adj = append(sq3_2.adj, &sq2_2)

	result := SearchMaze(&badMaze)
	if result {
		t.Errorf("BadMaze: Expected SearchMaze to return false but returned true")
	}

	goodMaze := Square{location: Coord{1, 1}}
	sq1_2 = Square{location: Coord{1, 2}}
	sq2_2 = Square{location: Coord{2, 2}}
	sq2_3 = Square{location: Coord{2, 3}}
	sq3_2 = Square{location: Coord{3, 2}}
	sq3_1 := Square{location: Coord{3, 1}, finish: true}

	goodMaze.adj = append(goodMaze.adj, &sq1_2)
	sq1_2.adj = append(sq1_2.adj, &badMaze)
	sq1_2.adj = append(sq1_2.adj, &sq2_2)
	sq2_2.adj = append(sq2_2.adj, &sq1_2)
	sq2_2.adj = append(sq2_2.adj, &sq2_3)
	sq2_2.adj = append(sq2_2.adj, &sq3_2)
	sq2_3.adj = append(sq2_3.adj, &sq2_2)
	sq3_2.adj = append(sq3_2.adj, &sq2_2)
	sq3_2.adj = append(sq3_2.adj, &sq3_1)
	sq3_1.adj = append(sq3_1.adj, &sq3_2)

	result = SearchMaze(&goodMaze)
	if !result {
		t.Errorf("GoodMaze: Expected SearchMaze to return true but returned false")
	}
}
