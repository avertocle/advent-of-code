package intz

import "fmt"

func Init3D(d1, d2, d3, val int) [][][]int {
	ans := make([][][]int, d1)
	for i := 0; i < d1; i++ {
		ans[i] = make([][]int, d2)
		for j := 0; j < d2; j++ {
			ans[i][j] = make([]int, d2)
			for k := 0; k < d3; k++ {
				ans[i][j][k] = val
			}
		}
	}
	return ans
}

func Count3d(arr3d [][][]int, val int) int {
	ctr := 0
	for _, arr2d := range arr3d {
		for _, arr1d := range arr2d {
			for _, v := range arr1d {
				if v == val {
					ctr++
				}
			}
		}
	}
	return ctr
}

/*
SetSub3D
sets a portion of a 3D array to a certain value
d1,d2,d3 are 2 elem 1d array having start and ends of the 3 dims
ignores areas of subset that are outside main set
returns count of values that were set
*/
func SetSub3D(arr3D [][][]int, d1, d2, d3 []int, val int) int {
	ctr := 0
	for i := d1[0]; i <= d1[1]; i++ {
		for j := d2[0]; j <= d2[1]; j++ {
			for k := d3[0]; k <= d3[1]; k++ {
				if InBounds3D([]int{i, j, k},
					[][]int{{0, len(arr3D)}, {0, len(arr3D[0])}, {0, len(arr3D[0][0])}}) {
					arr3D[i][j][k] = val
					ctr++
				}
			}
		}
	}
	//x := (d1[1] - d1[0]) * (d2[1] - d2[0]) * (d3[1] - d3[0])
	//if ctr != x {
	//	fmt.Printf("error intz.SetSub3D : incorrect count : got(%v) vs exp(%v)\n", ctr, x)
	//}
	return ctr
}

/*
InBounds3D
x : {i,j,k}
bounds : [[i_start,i_end],[..],[..]]
*/
func InBounds3D(x []int, bounds [][]int) bool {
	return (x[0] >= bounds[0][0] &&
		x[0] >= bounds[1][0] &&
		x[1] >= bounds[2][0] &&
		x[1] <= bounds[0][1] &&
		x[2] <= bounds[1][1] &&
		x[2] <= bounds[2][1])
}

func PPrint3D(arr [][][]int) {
	for _, d1 := range arr {
		PPrint2D(d1)
		fmt.Println()
	}
	fmt.Println()
}
