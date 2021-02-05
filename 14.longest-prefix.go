/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (36.30%)
 * Likes:    1444
 * Dislikes: 0
 * Total Accepted:    442.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 0
 * strs[i] 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 看了一下其他的解法, 这个叫纵向扫描, 应该是最快的, 横向扫描有缺陷, 若最后一个串为空串的话, 纵扫一趟就返回结果了, 横扫要len(strs) - 1趟才返回结果
	if strs == nil || len(strs) == 0 {
		return ""
	}

	// 以第一个串为蓝本, 不一定是最短的串, 所以下面的if需要先判断下标是否存在
	prefixLength := len(strs[0])

	for i := 0; i < prefixLength; i++ {
		for k := range strs {
			// 就是这里, 当strs[k]的长度等于i时, 是木有这个字符的, 直接返回即可
			if len(strs[k]) == i || strs[k][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	// 全遍历完了, 那蓝本就是最长公共前缀
	return strs[0]
}

// @lc code=end

