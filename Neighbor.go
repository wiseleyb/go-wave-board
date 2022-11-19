package main

type Neighbor struct {
	pos  string // up/down/left/right relative to cell. So "up" here would be the "down" side of cell above
	tile Tile
}

func newNeighbor(pos string, tile Tile) Neighbor {
	return Neighbor{pos: pos, tile: tile}
}
