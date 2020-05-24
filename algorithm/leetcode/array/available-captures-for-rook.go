// 999. 可以被一步捕获的棋子数
// https://leetcode-cn.com/problems/available-captures-for-rook/
package array

func numRookCaptures(board [][]byte) int {
	n := len(board)
	x, y := findRook(board)

	var result int
	// 左左
	for yi := y - 1; yi >= 0; yi-- {
		if board[x][yi] == 'p' {
			result++
			break
		}
		if board[x][yi] == 'B' {
			break
		}
	}

	// 右右
	for yi := y + 1; yi < n; yi++ {
		if board[x][yi] == 'p' {
			result++
			break
		}
		if board[x][yi] == 'B' {
			break
		}
	}

	// 上上
	for xi := x - 1; xi >= 0; xi-- {
		if board[xi][y] == 'p' {
			result++
			break
		}
		if board[xi][y] == 'B' {
			break
		}
	}

	// 下下
	for xi := x + 1; xi < n; xi++ {
		if board[xi][y] == 'p' {
			result++
			break
		}
		if board[xi][y] == 'B' {
			break
		}
	}

	return result
}
func findRook(board [][]byte) (int, int) {
	for i, arr := range board {
		for y, v := range arr {
			if v == 'R' {
				return i, y
			}
		}
	}
	panic("invalid input")
}
