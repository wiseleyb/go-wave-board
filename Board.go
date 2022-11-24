package main

import (
	"fmt"
	"sort"
)

type Board struct {
	cols        int
	rows        int
	grid        [][]Tile
	scores      [][]int
	tiles       Tiles
	tileWidth   int
	tileHeight  int
	tileSize    int
	boardWidth  int
	boardHeight int
	matches     map[int][]MatchSet
}

func newBoard(cols int, rows int, tiles Tiles) Board {
	b := Board{
		cols:        cols,
		rows:        rows,
		grid:        [][]Tile{},
		scores:      [][]int{},
		tiles:       tiles,
		tileWidth:   0,
		tileHeight:  0,
		tileSize:    0,
		boardWidth:  0,
		boardHeight: 0,
		matches:     map[int][]MatchSet{},
	}
	b.cols = cols
	b.rows = rows
	b.grid = make([][]Tile, b.rows)
	b.scores = make([][]int, b.rows)
	for i := range b.grid {
		b.grid[i] = make([]Tile, b.cols)
		b.scores[i] = make([]int, b.cols)
	}
	b.tiles = tiles
	b.tileWidth = tiles.tiles[0].cols
	b.tileHeight = tiles.tiles[0].rows
	b.tileSize = tiles.tiles[0].cols
	b.boardWidth = b.cols * b.tileWidth
	b.boardHeight = b.rows * b.tileHeight
	b.initBoard()
	b.run()
	return b
}

func (b *Board) run() {
	for {
		if b.applyMatch() {
			//b.print()
			//BoardToPng(b)
			return
		}
	}
}

func (b *Board) initBoard() {
	t := b.tiles.rndTile()
	c := newRnd(b.cols)
	r := newRnd(b.rows)
	for rowIdx, _ := range b.grid {
		for colIdx, _ := range b.grid[rowIdx] {
			b.grid[rowIdx][colIdx] = b.tiles.nullTile
		}
	}
	b.grid[r][c] = t
}

func (b *Board) scoreBoard() {
	b.matches = make(map[int][]MatchSet)
	for rowIdx, _ := range b.scores {
		for colIdx, _ := range b.scores[rowIdx] {
			t := b.grid[rowIdx][colIdx]
			if t.slug == "null" {
				// get sides for u/d/l/r
				score := 0
				foundTiles := make([]Tile, 0)
				neighbors := b.neighbors(rowIdx, colIdx)
				if len(neighbors) > 0 {
					for _, tile := range b.tiles.tiles {
						matches := 0
						for _, neighbor := range neighbors {
							if tile.edgeMatch(neighbor.pos, neighbor.tile) {
								matches += 1
							}
						}
						if matches == len(neighbors) {
							score += 1
							foundTiles = append(foundTiles, tile)
						}
					}
				}
				b.scores[rowIdx][colIdx] = score //len(b.neighbors(rowIdx, colIdx))
				b.matches[score] = append(b.matches[score], newMatchSet(rowIdx, colIdx, foundTiles))
			} else {
				b.scores[rowIdx][colIdx] = 0
			}
		}
	}
}

func (b *Board) applyMatch() bool {
	b.scoreBoard()
	//b.printScores()
	keys := make([]int, 0, len(b.matches))
	for k := range b.matches {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if len(keys) < 1 {
		return true
	}
	key := keys[len(keys)-1]
	if key < 1 {
		return true
	}
	setCnt := len(b.matches[key])
	matchSet := b.matches[key][newRnd(setCnt)]
	tileCnt := len(matchSet.tiles)
	if tileCnt > 0 {
		b.grid[matchSet.row][matchSet.col] = matchSet.tiles[newRnd(len(matchSet.tiles))]
		//b.print()
		return false
	} else {
		fmt.Println("done")
		return true
	}
}

func (b Board) neighbors(row int, col int) []Neighbor {
	var res []Neighbor
	// up
	if row > 0 {
		t := b.grid[row-1][col]
		if !(t.slug == "null") {
			res = append(res, newNeighbor("up", t))
		}
	}
	// down
	if row < b.rows-1 {
		t := b.grid[row+1][col]
		if !(t.slug == "null") {
			res = append(res, newNeighbor("down", t))
		}
	}
	// left
	if col > 0 {
		t := b.grid[row][col-1]
		if !(t.slug == "null") {
			res = append(res, newNeighbor("left", t))
		}
	}
	// right
	if col < b.cols-1 {
		t := b.grid[row][col+1]
		if !(t.slug == "null") {
			res = append(res, newNeighbor("right", t))
		}
	}
	return res
}

func (b Board) print() {
	for rowIdx, _ := range b.grid {
		for lineIdx, _ := range b.grid[rowIdx][0].tile {
			fmt.Print(rowIdx, " ")
			for _, tile := range b.grid[rowIdx] {
				for _, cell := range tile.tile[lineIdx] {
					fmt.Print(cell)
				}
			}
			fmt.Println("")
			/*
				if b.grid[rowIdx][colIdx].slug == b.tiles.nullTile.slug {
					fmt.Print(".")
				} else {
					fmt.Print("X")
				}
			*/
		}
		//fmt.Println("")
	}
}

func (b Board) printScores() {
	for rowIdx, _ := range b.scores {
		fmt.Print(rowIdx, " ")
		for _, val := range b.scores[rowIdx] {
			fmt.Print(val)
		}
		fmt.Println("")
	}
}

func newTdyBoard(width int, height int) Board {
	// TidByt is 62x32
	// 3x3 of that is 21x10
	return newBoard(width, height, newTilesTetris().tiles())
}

func newBoardFromYaml(width int, height int, yamlFileName string) Board {
	board := newBoard(width, height, yamlToTiles(yamlFileName))	
	fmt.Println(board.tiles)
	return board
}
