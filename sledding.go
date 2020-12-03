package main

type tree bool

type treeMap struct {
	grid       [][]tree
	currentX   int
	currentY   int
	simulatedX int
}

func makeTree(j rune) tree {
	if j == '#' {
		return tree(true)
	}

	return tree(false)
}

func buildTreeMap(input []string) *treeMap {
	var unfinishedMap [][]tree
	newRow := make([]tree, len(input[0]))
	unfinishedMap = make([][]tree, len(input))

	for irow, row := range input {
		for icol, column := range row {
			newRow[icol] = makeTree(column)
		}
		unfinishedMap[irow] = newRow
		newRow = make([]tree, len(row))
	}

	return &treeMap{unfinishedMap, 0, 0, 0}
}

func moveX(slope *treeMap, dx int) {
	slope.simulatedX += dx

	if (slope.currentX + dx) < len(slope.grid[0]) {
		slope.currentX += dx
		return
	}

	modifiedX := (slope.currentX + dx) % len(slope.grid[0])
	slope.currentX = modifiedX
}

func moveY(slope *treeMap, dy int) {
	if (slope.currentY + dy) < len(slope.grid) {
		slope.currentY += dy
		return
	}

	slope.currentY = len(slope.grid) - 1
}

func sledDown(slope *treeMap, dx int, dy int) {
	moveX(slope, dx)
	moveY(slope, dy)
}

func onTree(slope *treeMap) bool {
	return slope.grid[slope.currentY][slope.currentX] == tree(true)
}

func hitBottom(slope *treeMap) bool {
	return slope.currentY >= (len(slope.grid) - 1)
}

func resetSlope(slope *treeMap) {
	slope.currentX = 0
	slope.currentY = 0
	slope.simulatedX = 0
}

func findTreesOnPath(slope *treeMap, dx int, dy int) int {
	treesHit := 0

	for !hitBottom(slope) {
		sledDown(slope, dx, dy)
		if onTree(slope) {
			treesHit++
		}
	}

	return treesHit
}
