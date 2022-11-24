package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Tile struct {
	tile [][]string
	slug string // a short cut for this tile. Example 'up', 'down', 'left', 'right'
	rows int
	cols int
}

func newTile(tile [][]string, slug string) Tile {
	t := Tile{}
	t.tile = tile
	t.slug = slug
	t.rows = len(tile)
	t.cols = len(tile[0])
	return t
}

func newTileUniform(dimensions int, char string, name string) Tile {
	var res []string
	for i := 0; i < dimensions; i++ {
		var line bytes.Buffer
		for j := 0; j < dimensions; j++ {
			line.WriteString(char)
		}
		res = append(res, line.String())
	}
	stile := strings.Join(res, "\n")
	return newTile(stringToArr(stile), name)
}

func (t Tile) edgeMatch(direction string, tile Tile) bool {
	switch direction {
	case "up":
		if t.edge(direction) == tile.edge("down") {
			return true
		}
	case "down":
		if t.edge(direction) == tile.edge("up") {
			return true
		}
	case "left":
		if t.edge(direction) == tile.edge("right") {
			return true
		}
	case "right":
		if t.edge(direction) == tile.edge("left") {
			return true
		}
	}
	return false
}

func (t Tile) edge(direction string) string {
	var sb strings.Builder

	switch direction {
	case "up":
		{
			for i := 0; i < t.cols; i++ {
				sb.WriteString(t.tile[0][i])
			}
			return sb.String()
		}
	case "down":
		{
			for i := 0; i < t.cols; i++ {
				sb.WriteString(t.tile[t.rows-1][i])
			}
			return sb.String()
		}
	case "left":
		{
			for i := 0; i < t.rows; i++ {
				sb.WriteString(t.tile[i][0])
			}
			return sb.String()
		}
	case "right":
		{
			for i := 0; i < t.rows; i++ {
				sb.WriteString(t.tile[i][t.cols-1])
			}
			return sb.String()
		}
	}
	return "nope"
}

func (t Tile) print() {
	fmt.Println("---")
	fmt.Println("slug:", t.slug)
	fmt.Println("rows:", t.rows)
	fmt.Println("cols:", t.cols)
	for _, row := range t.tile {
		fmt.Println(row)
	}
	fmt.Println("")
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
			fmt.Println(row, col)
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
