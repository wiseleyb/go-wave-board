package main

import (
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
