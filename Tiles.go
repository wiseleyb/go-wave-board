package main

type Tiles struct {
	nullTile Tile
	tiles    []Tile
}

func newTiles(tiles []Tile) Tiles {
	t := Tiles{}
	t.tiles = tiles
	t.nullTile = newTileUniform(len(t.tiles[0].tile[0]), ".", "null")
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
