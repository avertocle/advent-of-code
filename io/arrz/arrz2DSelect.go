package arrz

import "github.com/avertocle/contests/io/tpz"

func IsValidIndexCriterion2D[T tpz.Primitive](grid [][]T, idx []int) bool {
	return idx[0] >= 0 && idx[0] < len(grid) && idx[1] >= 0 && idx[1] < len(grid[0])
}

func MakeValueCriterion2D[T tpz.Primitive](val T) CriterionFunc[T] {
	return func(grid [][]T, index []int) bool {
		return grid[index[0]][index[1]] == val
	}
}

func GenericSelect2D[T tpz.Primitive](grid [][]T, indexList [][]int, criteria []CriterionFunc[T]) [][]int {
	ans := make([][]int, 0)
	if indexList != nil {
		for _, idx := range indexList {
			if matchAllCriteria(grid, idx, criteria) {
				ans = append(ans, idx)
			}
		}
	} else {
		for i, row := range grid {
			for j, _ := range row {
				if matchAllCriteria(grid, []int{i, j}, criteria) {
					ans = append(ans, []int{i, j})
				}
			}
		}
	}
	return ans
}

func matchAllCriteria[T tpz.Primitive](grid [][]T, index []int, criteria []CriterionFunc[T]) bool {
	for _, criterion := range criteria {
		if !criterion(grid, index) {
			return false
		}
	}
	return true
}

/*
CriterionFunc
func(grid2D, index, value)
serves as a criteria function for both index and value based criteria
*/
type CriterionFunc[T tpz.Primitive] func([][]T, []int) bool

func Neighbours2D(index []int) [][]int {
	return [][]int{
		{index[0] - 1, index[1]},
		{index[0] + 1, index[1]},
		{index[0], index[1] - 1},
		{index[0], index[1] + 1},
	}
}
func Neighbours2DWithDiag(index []int) [][]int {
	return [][]int{
		{index[0] - 1, index[1] - 1},
		{index[0] - 1, index[1]},
		{index[0] - 1, index[1] + 1},
		{index[0], index[1] - 1},
		{index[0], index[1] + 1},
		{index[0] + 1, index[1] - 1},
		{index[0] + 1, index[1]},
		{index[0] + 1, index[1] + 1},
	}
}
