package main

type Tiles struct {
	nullTile Tile
	tiles    []Tile
}

func newTiles(tiles []Tile, nullTile Tile) Tiles {
	t := Tiles{}
	t.tiles = tiles
	t.nullTile = nullTile
	return t
}

func (t Tiles) rndTile() Tile {
	idx := newRnd(len(t.tiles))
	return t.tiles[idx]
}

/*
func (t Tiles) matches(direction string, tile Tile) []Tile {
	var res = []Tile{}
	side := tile.edge(direction)
	switch direction {
	case "up":
		{
		}
	case "down":
		{
		}
	case "left":
		{
		}
	case "right":
		{
		}
	}
	return res
}
*/
