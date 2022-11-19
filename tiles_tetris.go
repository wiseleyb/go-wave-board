package main

import (
	"strings"
)

type TilesTetris struct {
	m map[string]Tile
}

func (t TilesTetris) tiles() Tiles {
	tt := Tiles{}
	tt.tiles = append(tt.tiles,
		t.blank(),
		t.up(),
		t.down(),
		t.left(),
		t.right())

	tt.nullTile = t.null()

	return tt
}

func newTilesTetris() TilesTetris {
	tt := TilesTetris{}
	tt.m = make(map[string]Tile)
	tt.m["null"] = tt.null()
	tt.m["blank"] = tt.blank()
	tt.m["up"] = tt.up()
	tt.m["down"] = tt.down()
	tt.m["left"] = tt.left()
	tt.m["right"] = tt.right()
	return tt
}

func (t TilesTetris) null() Tile {
	s := `...
		  ...
		  ...`
	return newTile(stringToArr(s), "null")
}

func (t TilesTetris) blank() Tile {
	s := `___
		  ___ 
		  ___`
	return newTile(stringToArr(s), "blank")
}

func (t TilesTetris) up() Tile {
	/*
		s := `abc
		      def
			  ghi`
	*/
	s := `_X_
	      XXX
		  ___`
	return newTile(stringToArr(s), "up")
}

func (t TilesTetris) down() Tile {
	s := `___
		  XXX
		  _X_`
	return newTile(stringToArr(s), "down")
}

func (t TilesTetris) left() Tile {
	s := `_X_
		  XX_
		  _X_`
	return newTile(stringToArr(s), "left")
}

func (t TilesTetris) right() Tile {
	s := `_X_
		  _XX
		  _X_`
	return newTile(stringToArr(s), "right")
}

func stringToArr(str string) [][]string {
	arr := strings.Split(str, "\n")
	cols := len(strings.TrimSpace(arr[0]))
	rows := len(arr)
	res := make([][]string, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]string, rows)
	}
	for row := 0; row < rows; row++ {
		rowstr := strings.Split(strings.TrimSpace(arr[row]), "")
		//fmt.Println(rowstr)
		for col := 0; col < cols; col++ {
			res[row][col] = rowstr[col]
		}
	}
	// fmt.Println(res)
	return res
}
