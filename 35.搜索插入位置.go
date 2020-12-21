/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.23%)
 * Likes:    770
 * Dislikes: 0
 * Total Accepted:    292.8K
 * Total Submissions: 624.2K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 * 
 * 你可以假设数组中无重复元素。
 * 
 * 示例 1:
 * 
 * 输入: [1,3,5,6], 5
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: [1,3,5,6], 2
 * 输出: 1
 * 
 * 
 * 示例 3:
 * 
 * 输入: [1,3,5,6], 7
 * 输出: 4
 * 
 * 
 * 示例 4:
 * 
 * 输入: [1,3,5,6], 0
 * 输出: 0
 * 
 * 
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	/* 这题没什么好说的, 遍历即可
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}

	// 若是执行到此, 只有一种情况, 数组中没有target, 同时数组的最后一个数是小于target的, 此时target插入的位置就是len(nums)
	return len(nums)
	*/

	// 你个笨比, 看见排序数组你的第一想法居然不是二分法, 真是笨比, 还搁着没什么好说的遍历即可, O(logN)不必O(N)快吗
	// 此题分两种情况, target存在, 没什么好说的, 正统二分即可; 不存在, 那到最后的情况一定是nums[R] < target < nums[L], L == R + 1, 此时返回L即可
	L, R := 0, len(nums) - 1
	for L <= R {
		mid := L + (R - L) >> 1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	return L
}
// @lc code=end

