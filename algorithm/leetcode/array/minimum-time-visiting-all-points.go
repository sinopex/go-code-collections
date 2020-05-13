// 1266. 访问所有点的最小时间
// https://leetcode-cn.com/problems/minimum-time-visiting-all-points/
package array

func minTimeToVisitAllPoints(points [][]int) int {
	totalStep := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	step := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for i := 0; i < len(points)-1; i++ {
		totalStep += max(step(points[i][1], points[i+1][1]), step(points[i][0], points[i+1][0]))
	}

	return totalStep
}
