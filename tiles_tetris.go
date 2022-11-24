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
		t.right(),
		t.cornerBottomLeft(),
		t.cornerBottomRight(),
		t.cornerTopLeft(),
		t.cornerTopRight(),
		t.horizontal(),
		t.vertical(),
		t.endUp(),
		t.endDown(),
		t.endLeft(),
		t.endRight(),
		t.cross())

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
	tt.m["cornerBottomLeft"] = tt.cornerBottomLeft()
	tt.m["cornerBottomRight"] = tt.cornerBottomRight()
	tt.m["cornerTopLeft"] = tt.cornerTopLeft()
	tt.m["cornerTopRight"] = tt.cornerTopRight()
	tt.m["horizontal"] = tt.horizontal()
	tt.m["vertical"] = tt.vertical()
	tt.m["endUp"] = tt.endUp()
	tt.m["endDown"] = tt.endDown()
	tt.m["endLeft"] = tt.endLeft()
	tt.m["endRight"] = tt.endRight()
	tt.m["cross"] = tt.cross()
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

func (t TilesTetris) cornerTopLeft() Tile {
	s := `_X_
		  XX_
		  ___`
	return newTile(stringToArr(s), "corner-top-left")
}

func (t TilesTetris) cornerTopRight() Tile {
	s := `_X_
		  _XX
		  ___`
	return newTile(stringToArr(s), "corner-top-right")
}

func (t TilesTetris) cornerBottomLeft() Tile {
	s := `___
		  XX_
		  _X_`
	return newTile(stringToArr(s), "corner-bottom-left")
}

func (t TilesTetris) cornerBottomRight() Tile {
	s := `___
		  _XX
		  _X_`
	return newTile(stringToArr(s), "corner-bottom-right")
}

func (t TilesTetris) horizontal() Tile {
	s := `___
		  XXX
		  ___`
	return newTile(stringToArr(s), "horizontal")
}

func (t TilesTetris) vertical() Tile {
	s := `_X_
		  _X_
		  _X_`
	return newTile(stringToArr(s), "vertical")
}

func (t TilesTetris) endUp() Tile {
	s := `_X_
		  _X_
		  ___`
	return newTile(stringToArr(s), "endUp")
}

func (t TilesTetris) endDown() Tile {
	s := `___
		  _X_
		  _X_`
	return newTile(stringToArr(s), "endDown")
}

func (t TilesTetris) endLeft() Tile {
	s := `___
		  XX_
		  ___`
	return newTile(stringToArr(s), "endLeft")
}

func (t TilesTetris) endRight() Tile {
	s := `___
		  _XX
		  ___`
	return newTile(stringToArr(s), "endRight")
}

func (t TilesTetris) cross() Tile {
	s := `_X_
		  XXX
		  _X_`
	return newTile(stringToArr(s), "cross")
}

func stringToArr(str string) [][]string {
	arr := jsondata(strings.Split(str, "\n"))
	arr.delete("")
	cols := len(strings.TrimSpace(arr[0]))
	rows := len(arr)
	res := make([][]string, cols)
	//fmt.Println(arr, cols, rows)
	for i := 0; i < cols; i++ {
		res[i] = make([]string, rows)
	}
	for row := 0; row < rows; row++ {
		rowstr := jsondata(strings.Split(strings.TrimSpace(arr[row]), ""))
		rowstr.delete("")
		//fmt.Println(rowstr)
		for col := 0; col < cols; col++ {
			res[row][col] = rowstr[col]
		}
	}
	// fmt.Println(res)
	return res
}

type jsondata []string

func (j *jsondata) delete(selector string) {
	var r jsondata
	for _, str := range *j {
		if str != selector {
			r = append(r, str)
		}
	}
	*j = r
}
