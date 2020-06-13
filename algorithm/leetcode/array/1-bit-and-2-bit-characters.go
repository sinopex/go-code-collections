// 717. 1比特与2比特字符
// https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
package array

func isOneBitCharacter(bits []int) bool {
	if len(bits) <= 1 {
		return true
	}
	counter := 0
	for i := len(bits) - 2; i > -1; i-- {
		if bits[i] == 0 {
			break
		}
		counter++
	}
	return counter&1 == 0
}
func isOneBitCharacter2(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits)-1
}

func isOneBitCharacter1(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	// 如果只有1位就可以定胜负了
	if len(bits) == 1 {
		if bits[0] == 0 {
			return true
		} else {
			return false
		}
	}
	return bits[0] == 0 && isOneBitCharacter1(bits[1:]) ||
		bits[0] == 1 && bits[1] == 0 && isOneBitCharacter1(bits[2:]) ||
		bits[0] == 1 && bits[1] == 1 && isOneBitCharacter1(bits[2:])
}
