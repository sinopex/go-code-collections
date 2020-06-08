// 1275. 找出井字棋的获胜者
// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game/
package array

func tictactoe(moves [][]int) string {
	const (
		StatusEmpty = iota
		StatusPending
		StatusDraw
	)
	type line struct {
		status int
		flag   string
		count  int
	}

	// 一共存在8条直线，每条直线先初始化
	lines := [8]line{
		0: {status: StatusEmpty}, // 第一行
		1: {status: StatusEmpty}, // 第二行
		2: {status: StatusEmpty}, // 第三行
		3: {status: StatusEmpty}, // 第一列
		4: {status: StatusEmpty}, // 第二列
		5: {status: StatusEmpty}, // 第三列
		6: {status: StatusEmpty}, // 左上到右下
		7: {status: StatusEmpty}, // 右上到左下
	}

	// 每个坐标点存在的直线列表
	pointerAtLines := map[int][]int{
		0: {0, 3, 6},
		1: {0, 4},
		2: {0, 5, 7},
		3: {1, 3},
		4: {1, 4, 6, 7},
		5: {1, 5},
		6: {2, 3, 7},
		7: {2, 4},
		8: {2, 5, 6},
	}

	flag := "A"
	count := 0
	for i := 0; i < len(moves); i++ {
		count++
		pointer := moves[i][0]*3 + moves[i][1]
		for _, lineId := range pointerAtLines[pointer] {
			// 计算当前线条的可能性
			if lines[lineId].status == StatusEmpty {
				lines[lineId].flag = flag
				lines[lineId].status = StatusPending
				lines[lineId].count++
			} else if lines[lineId].status == StatusPending {
				if lines[lineId].flag != flag {
					lines[lineId].status = StatusDraw
				} else {
					lines[lineId].count++
					if lines[lineId].count == 3 {
						return flag
					}
				}
			} else {

			}
		}

		// 如果没有棋子可走
		if count == 9 {
			return "Draw"
		}

		// A和B交手
		if flag == "A" {
			flag = "B"
		} else {
			flag = "A"
		}
	}

	return "Pending"
}
