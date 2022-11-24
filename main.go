package main

func main() {
	GrnTest()
	/*
		fname := "tmp/wave.png"
		//board := newBoardFromYaml(20, 20, "tile-sets/ts-tetris.yaml")
		//board := newBoardFromYaml(20, 20, "tile-sets/ts-4x4.yaml")
		//board := newBoardFromYaml(20, 20, "tile-sets/ts-tetris2.yaml")
		board := newBoardFromYaml(20, 20, "tile-sets/ts-5x5.yaml")
		board.print()
		fmt.Println(board.tiles)
		br := newBoardRenderer(board, 4)
		br.boardToPng()
		br.save(fname)
		board.print()
		fmt.Println(board.tileSize)
	*/
	/*
		board := newTdyBoard(20, 20)
		board.print()
		br := newBoardRenderer(board, 2)
		br.boardToPng()
		br.save()
	*/

	//b := newTdyBoard()
	//b.run()
	//drawTest()

	/*
		tt := newTilesTetris()
		b := newBoardRenderer(100, 100)
		b.drawTile(tt.up(), 1, 1)
		b.drawTile(tt.down(), 2, 2)
		b.save()
	*/

	//SvgTest()
	/*
		b.print()
		for i := 1; i < 30; i++ {
			b.printScores()
			done := b.applyMatch()
			if done == true {
				continue
			} else {
				b.scoreBoard()
			}
		}
	*/
	/*
		fmt.Println(b)
		t := b.tiles.tiles[3]
		t.print()
		for _, tile := range b.tiles.tiles {
			fmt.Println(tile.slug, t.edgeMatch("up", tile))
			//tile.print()
			//fmt.Println("")
		}
		fmt.Println(b.grid)

		tt := newTilesTetris()
		tt.print("null")
		tt.print("blank")
		tt.print("up")
		tt.print("down")
		tt.print("left")
		tt.print("right")
		fmt.Println("top edge of up tile", tt.up().edge("up"))
		fmt.Println("bottom edge of up tile", tt.up().edge("down"))
		fmt.Println("left edge of up tile", tt.up().edge("left"))
		fmt.Println("right edge of up tile", tt.up().edge("right"))

		fmt.Println("tiles", tt.tiles())
	*/
	//tile := newTileBlank()
	//tile.PrintT()

	//tiles := newTiles()
	//tiles.PrintT()

	/*
		grid := newGrid()
		printGrid(grid)
		grid.scoreMatches()
		fmt.Println(grid.matches[0][0])
		fmt.Println(grid.matches)
	*/

	//testSides()

	/*
		x := 3
		arr := make([]int, x)
		fmt.Println(arr)
		s := "{x}"
		fmt.Println(s)

		var sb strings.Builder
		sb.WriteString("a")
		fmt.Println(sb.String())
	*/
}
