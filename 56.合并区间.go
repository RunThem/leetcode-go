/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.44%)
 * Likes:    731
 * Dislikes: 0
 * Total Accepted:    174.2K
 * Total Submissions: 398.2K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: intervals = [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 * 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * intervals[i][0] <= intervals[i][1]
 * 
 * 
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 第一想法就是先排序, 以intervals[i][0]为基准
	// 那以下数据就变成了这样[[1,3], [2,6], [8,10], [15,18]]
	// 1, 2, 8,  10
	// 3, 6, 10, 18
	// 只要右上角的比左下角的小或等于, 将上面的两个值对调, 若右下角小于左下角的, 那也对调, 遍历下去, 当上面的两条规则不成立时, 一个区间就出来了, 就是intervals[i]
	// 

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i][1] >= intervals[i + 1][0] {
			if intervals[i][1] > intervals[i + 1][1] {
				intervals[i + 1][1] = intervals[i][1]
			}
			intervals[i + 1][0] = intervals[i][0]
			continue
		}
		ret = append(ret, intervals[i])
	}
	ret = append(ret, intervals[len(intervals) - 1])
	
	return ret
}
// @lc code=end

