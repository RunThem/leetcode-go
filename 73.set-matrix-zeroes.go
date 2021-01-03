/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode-cn.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (55.17%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    63.2K
 * Total Submissions: 113K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
 * 
 * 示例 1:
 * 
 * 输入: 
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * 输出: 
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: 
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * 输出: 
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 * 
 * 进阶:
 * 
 * 
 * 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个常数空间的解决方案吗？
 * 
 * 
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	/* 第一想法就是需要一个额外空间的辅助矩阵了, 但是不符合题意, 当时没想出好方法, 后面看了其他人的解法才明白
	M := len(matrix)
	N := len(matrix[0])
	if M == 0 {
		return
	}
	
	// 标记, 但是有额外的空间占用, 16ms, 6.4Mb
	MapRow := make(map[int]bool)
	MapCol := make(map[int]bool)
	
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] == 0 {
				if _, ok := MapRow[i]; !ok {
					MapRow[i] = true
				}
				if _, ok := MapCol[j]; !ok {
					MapCol[j] = true
				}
			}
		}
	}
	
	for k, _ := range MapRow {
		for i := 0; i < N; i++ {
			matrix[k][i] = 0
		}
	}
	
	for k, _ := range MapCol {
		for i := 0; i < M; i++ {
			matrix[i][k] = 0
		}
	} */

	// 第二种方法, 将其映射至第一行, 第一列
	M := len(matrix)
	N := len(matrix[0])

	isFristRowHaveZero := false
	isFristColHaveZero := false
	
	// 标记第一行和第一列是否需要清除
	for i := 0; i < N; i++ {
		if matrix[0][i] == 0 {
			isFristRowHaveZero = true
			break
		}
	}
	
	for i := 0; i < M; i++ {
		if matrix[i][0] == 0 {
			isFristColHaveZero = true
			break
		}
	}
	
	// 开始映射
	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	
	
	// 清除第一行第一列外的数据
	for i := 1; i < M; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < N; j++ {
				matrix[i][j] = 0
			}
		}
	}
	
	for i := 1; i < N; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < M; j++ {
				matrix[j][i] = 0
			}
		}
	}
	
	// 清除第一行第一列
	if isFristRowHaveZero {
		for i := 0; i < N; i++ {
			matrix[0][i] = 0
		}
	}
	
	if isFristColHaveZero {
		for i := 0; i < M; i++ {
			matrix[i][0] = 0
		}
	}
}
// @lc code=end

