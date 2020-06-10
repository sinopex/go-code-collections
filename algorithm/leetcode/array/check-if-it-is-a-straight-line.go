// 1232. 缀点成线
// https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
package array

func checkStraightLine(coordinates [][]int) bool {
	x0, y0 := coordinates[0][0]-coordinates[1][0], coordinates[0][1]-coordinates[1][1]

	for i := 1; i < len(coordinates)-1; i++ {
		xi, yi := coordinates[i][0]-coordinates[i+1][0], coordinates[i][1]-coordinates[i+1][1]
		if xi*y0 != yi*x0 {
			return false
		}
	}
	return true
}

func checkStraightLine1(coordinates [][]int) bool {
	ratio := getRatio(coordinates[0], coordinates[1])

	for i := 1; i < len(coordinates)-1; i++ {
		if getRatio(coordinates[i], coordinates[i+1]) != ratio {
			return false
		}
	}

	return true
}
func getRatio(pointerX, pointerY []int) float64 {
	px := pointerX[0] - pointerY[0]
	py := pointerX[1] - pointerY[1]

	if px == 0 {
		return 0
	}
	if py == 0 {
		return 1
	}

	return float64(px) / float64(py)
}
