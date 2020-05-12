// 657 机器人能否回到原点
// https://leetcode-cn.com/problems/robot-return-to-origin/
package string

func judgeCircle(moves string) bool {
	var set = map[rune]int{
		'U': 0,
		'L': 0,
		'R': 0,
		'D': 0,
	}
	for _, s := range moves {
		set[s]++
	}

	return (set['U']-set['D'] == 0) && (set['L']-set['R'] == 0)
}
