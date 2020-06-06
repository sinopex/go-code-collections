// 1184. 公交站间的距离
// https://leetcode-cn.com/problems/distance-between-bus-stops/
package array

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum, direct := 0, 0

	if start > destination {
		start, destination = destination, start
	}

	for i := 0; i < len(distance); i++ {
		sum += distance[i]
		if i >= start && i < destination {
			direct += distance[i]
		}
	}

	if sum-direct < direct {
		return sum - direct
	}

	return direct
}
