/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (39.59%)
 * Likes:    643
 * Dislikes: 0
 * Total Accepted:    267.3K
 * Total Submissions: 675.3K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 * 
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 * 
 * 示例 1:
 * 
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 * 
 * 
 * 说明:
 * 
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 * 
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 * 
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	/* 简单点就是暴力求解了, 这没什么好说的, 复杂度为O((M-N)N)
	hLen := len(haystack)
	nLen := len(needle)

	if nLen == 0 {
		return 0
	}
	
	for i := 0; i <= hLen - nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
		
	return -1
	*/

	/* 再去找字符串匹配的算法, 那第一眼就看见的就是KMP算法了, 是有点复杂的, next数组(就是前缀表(prefix table)定制版本)还是有点烧脑的
	前缀表是记录下标i之前(包括i)的字符串中, 有多大长度的相同前缀后缀
	它的任务就是当当前位置匹配失败, 找到之前已经匹配上的位置, 再重新匹配, 意味着在某个字符失配时, 前缀表会告诉你下一步匹配中, 模式串应该跳到哪个位置
	一句话, next[i]中的值就是模式串中下标i之前(不包括i)的字符串中, 有多大长度的相同前缀后缀. 看看, 是不是与上面的前缀表定义不一样, 其实很好理解为什么是不包含i的字符串, 因为是当i指向的字符失配时才利用next数组去重新匹配, 既然i指向的字符失配, 自然就是i之前的字符串的最大长度相同前缀后缀了, 此时就不能包含i, i之前的字符串的最大长度相同前缀后缀就是prefix_table[i-1]的值了, 这就是prefix_table整体右移一个的原因
	那此时就很好理解了, next = {next[0] = -1, next[i] = prefix_table[i-1]}, next[0]是当prefix_table右移时多的一个, 要区别开最大长度相同前缀后缀为0的情况
	当然, 这只是KMP的不同实现, 我并不采用这种定制的前缀表, 我选择使next == prefix_table, 这样好理解点
	https://www.zhihu.com/question/21923021/answer/1032665486
	*/

	if len(needle) == 0 {
		return 0
	}

	return KMPSearch(haystack, needle)
}

func KMPSearch(haystack string, needle string) int {
	tar, pos := 0, 0
	next := prefix_table(needle)
	
	for tar < len(haystack) {
		if haystack[tar] == needle[pos] {
			tar++
			pos++
		} else if pos != 0 {
			pos = next[pos - 1]
		} else {
			tar++
		}

		if pos == len(needle) {
			return tar - pos
		}
	}

	return -1
}

func prefix_table(s string) []int {
	next := make([]int, len(s))
	
	for i, j := 1, 0; i < len(s); {
		if s[i] == s[j] {
			j++
			next[i] = j
			i++
		} else if j != 0 {
			j = next[j-1]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
// @lc code=end

