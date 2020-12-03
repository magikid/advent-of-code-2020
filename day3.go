package main

type tree bool

type treeMap struct {
	grid     [][]tree
	currentX int
	currentY int
}

func makeTree(j rune) tree {
	if j == '#' {
		return tree(true)
	}

	return tree(false)
}

func makeTreeMap(input []string) *treeMap {
	var unfinishedMap [][]tree
	newRow := make([]tree, len(input[0]))

	for _, row := range input {
		for icol, column := range row {
			newRow[icol] = makeTree(column)
		}
		unfinishedMap = append(unfinishedMap, newRow)
		newRow = make([]tree, len(row))
	}

	return &treeMap{unfinishedMap, 0, 0}
}

// Day3Solution1 finds the number of tress in a path on the slope right 3, down 1
func Day3Solution1(input []string, results chan string) {
	results <- ""
}

// Day3Solution2 WIP
func Day3Solution2(input []string, results chan string) {
	results <- ""
}
