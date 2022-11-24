package main

import (
	"bytes"
	"fmt"
	"strings"
)

type GenRndTile struct {
	dimensions int
	tileCount  int
	tiles      Tiles
	board      Board
}

func GrnTest() {
	grn := newGenRndTiles(20, 20, 3, 50)
	fmt.Println(grn.tiles)
	br := newBoardRenderer(grn.board, 4)
	br.boardToPng()
	br.save("tmp/wave.png")
	fmt.Println(grn.tiles)
	fmt.Println(grn.board)
	grn.board.printScores()
}

func newGenRndTiles(width int, height int, dimensions int, tileCount int) GenRndTile {
	grn := GenRndTile{dimensions: dimensions, tileCount: tileCount}
	for tc := 0; tc < tileCount; tc++ {
		var res []string
		for i := 0; i < grn.dimensions; i++ {
			var line bytes.Buffer
			for j := 0; j < grn.dimensions; j++ {
				if newRnd(2) == 0 {
					line.WriteString("_")
				} else {
					line.WriteString("X")
				}
			}
			res = append(res, line.String())
		}
		stile := strings.Join(res, "\n")
		tname := fmt.Sprintf("tile%d", tc)
		t := newTile(stringToArr(stile), tname)
		grn.tiles.tiles = append(grn.tiles.tiles, t)
	}
	grn.tiles.tiles = append(grn.tiles.tiles, newTileUniform(dimensions, "_", "blank"))
	grn.tiles.nullTile = newTileUniform(dimensions, ".", "null")
	grn.board = newBoard(width, height, grn.tiles)

	return grn
}
