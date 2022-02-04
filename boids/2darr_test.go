package main_test

import "testing"

func TestGame_2dArray(t *testing.T) {
	/* renders:
	[[-1 -1 -1] # 1st row
	 [-1 -1 -1] # 2nd row
	 [-1 -1 -1]] # 3rd row
	*/

	var m [3][3]int
	// iterate over rows aka first dimension
	for i, row := range m {
		// iterate over elements in each row aka 2nd dimension
		for j := range row {
			m[i][j] = -1
		}
	}

	t.Log(m) // output [[-1 -1 -1] [-1 -1 -1] [-1 -1 -1]]
}
