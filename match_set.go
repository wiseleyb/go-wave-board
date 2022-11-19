package main

type MatchSet struct {
	row   int
	col   int
	tiles []Tile
}

func newMatchSet(row int, col int, tiles []Tile) MatchSet {
	return MatchSet{row: row, col: col, tiles: tiles}
}
