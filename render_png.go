package main

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type BoardRenderer struct {
	colors       [7]color.RGBA
	defaultColor color.RGBA
	board        Board
	scale        int
	height       int
	width        int
	cc           *canvas.Canvas
	ctx          *canvas.Context
	tileSize     int
	tileWidth    int
}

func newBoardRenderer(board Board, scale int) BoardRenderer {
	height := board.rows * board.tileSize * scale
	width := board.cols * board.tileSize * scale
	c := canvas.New(float64(width), float64(height))
	b := BoardRenderer{height: height, width: width, cc: c, ctx: canvas.NewContext(c)}
	b.board = board
	b.scale = scale
	b.tileSize = board.tileSize
	b.tileWidth = scale

	b.colors[0] = canvas.Red
	b.colors[1] = canvas.Orange
	b.colors[2] = canvas.Yellow
	b.colors[3] = canvas.Green
	b.colors[4] = canvas.Blue
	b.colors[5] = canvas.Indigo
	b.colors[6] = canvas.Purple
	b.defaultColor = canvas.Lightblue

	return b
}

func (b BoardRenderer) save(fname string) {
	renderers.Write(fname, b.cc, canvas.DPMM(10.2))
}

func (b BoardRenderer) boardToPng() {
	fmt.Println("boardToPng", b.board.grid)
	for rowIdx, _ := range b.board.grid {
		for colIdx, tile := range b.board.grid[rowIdx] {
			fmt.Println("process cell: ", rowIdx, colIdx)
			b.drawTile(tile, rowIdx, colIdx)
		}
	}
}

func (b BoardRenderer) rndColor() color.RGBA {
	return b.colors[newRnd(6)]
}

func (b BoardRenderer) drawTile(t Tile, xint int, yint int) {
	fmt.Println("tile:", t.slug)
	t.print()

	if t.slug == "null" {
		b.ctx.SetFillColor(canvas.Lightgray)
	} else {
		//b.ctx.SetFillColor(canvas.Lightblue)
		//b.ctx.SetFillColor(b.rndColor())
		b.ctx.SetFillColor(b.defaultColor)
	}
	b.ctx.SetStrokeColor(canvas.Transparent)

	ts := float64(b.tileSize)
	tw := float64(b.tileWidth)
	baseX := float64(xint) * tw * ts
	baseY := float64(yint) * tw * ts

	for cellRowIdx, _ := range t.tile {
		for cellColIdx, cell := range t.tile[cellRowIdx] {
			if cell == "X" || t.slug == "null" {
				x := baseX + float64(cellRowIdx)*tw
				y := baseY + float64(cellColIdx)*tw
				rect := canvas.Rectangle(tw, tw)
				fmt.Println("pcell:", x, y, rect)
				b.ctx.DrawPath(x, y, rect)
			}
		}
	}
	b.ctx.SetFillColor(canvas.Transparent)
	b.ctx.SetStrokeColor(canvas.Black)
	rect := canvas.Rectangle(tw*ts, tw*ts)
	b.ctx.DrawPath(baseX, baseY, rect)
}
