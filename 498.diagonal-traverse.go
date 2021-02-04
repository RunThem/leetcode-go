/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 *
 * https://leetcode-cn.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (39.68%)
 * Likes:    168
 * Dislikes: 0
 * Total Accepted:    29.6K
 * Total Submissions: 69K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 
 * 输出:  [1,2,4,7,5,3,6,8,9]
 * 
 * 解释:
 * 
 * 
 * 
 * 
 * 
 * 说明:
 * 
 * 
 * 给定矩阵中的元素总数不会超过 100000 。
 * 
 * 
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
	// 校验参数
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	row, col, x, y := len(matrix), len(matrix[0]), 0, 0
	d, dir := 0, [2][2]int {
		{-1, 1},
		{1, -1},
	}
	total := row * col
	res := make([]int, total)

	// 当row或col为1时, 遍历为total次
	for i := 0; i < total; {
		// 每次都走完一条线
		for x >= 0 && x < row && y >= 0 && y < col {
			res[i] = matrix[x][y]
			i++
			x += dir[d][0]
			y += dir[d][1]
		}
		// 这里是最有意思的, d在0与1之间反复横跳, 有点东西
		d = (d + 1) % 2

		// 我之前的想法老师用两个if写完所有的逻辑, 无法处理两个特殊的点
		// 下面就是当走到边界时怎么刷新x, y的值, 有两个点特殊, [0][col - 1]与[row - 1][0], 这两个点会命中两个if, 其他都是只命中一个if
		if x == row {
			x--
			y += 2
		}
		if y == col {
			y--
			x += 2
		}
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
	}
	return res	
}
// @lc code=end

