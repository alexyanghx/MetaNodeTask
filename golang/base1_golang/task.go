package base1_golang

import (
	"fmt"
	"strconv"
)

// 136. 只出现一次的数字：给定一个非空整数数组，
// 除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
// 结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func Task1() {
	method := func(nums []int) int {
		numCount := make(map[int]int)
		for _, num := range nums {
			numCount[num]++
		}

		for num, count := range numCount {
			if count == 1 {
				return num
			}
		}

		panic("no unique element found")
	}
	nums := []int{2, 2, 1}
	fmt.Println(method(nums))
}

// 回文数,
// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数
func Task2() {
	method := func(x int) bool {
		str := strconv.Itoa(x)
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	x := 121
	fmt.Println(method(x)) // 输出 true
}

// 有效的括号 ,
// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func Task3() {
	method := func(s string) bool {
		stack := []rune{}
		mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

		for _, char := range s {
			if char == '(' || char == '{' || char == '[' {
				stack = append(stack, char)
			} else if char == ')' || char == '}' || char == ']' {
				if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
		return len(stack) == 0
	}

	str := "({[]})"
	fmt.Println(method(str)) // 输出 true
}

// 最长公共前缀,
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀
func Task4() {
	method := func(strs []string) string {
		if len(strs) == 0 {
			return ""
		}

		prefix := strs[0]
		if len(prefix) == 0 {
			return ""
		}

		for i := 1; i < len(strs); i++ {
			j := 0
			for ; j < min(len(strs[i]), len(prefix)); j++ {
				if strs[i][j] != prefix[j] {

					break
				}
			}
			prefix = prefix[:j]
			if len(prefix) == 0 {
				return ""
			}
		}
		return prefix
	}

	fmt.Println(method([]string{"dog", "racecar", "car"}))
}

// 删除排序数组中的重复项 ,
// 考察：数组/切片操作
// 题目：给定一个排序数组，你需要在原地删除重复出现的元素
func Task5() {
	method := func(nums []int) []int {
		j := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[j] {
				j++
				nums[j] = nums[i]
				fmt.Println(nums)
			}
		}

		return nums[:j+1]
	}

	nums := []int{1, 1, 2, 2, 3}

	fmt.Println(method(nums)) // 输出完成提示
}

// 加一 ,
// 难度：简单
// 考察：数组操作、进位处理
// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func Task6() {
	method := func(digits []int) []int {
		i := len(digits) - 1
		for ; i >= 0; i-- {
			if digits[i] < 9 {
				digits[i]++
				return digits
			}

			digits[i] = 0

		}
		if i < 0 {
			digits = append([]int{1}, digits...)
		}

		return digits

	}

	digits := []int{9, 9, 9}
	fmt.Println(method(digits)) // 输出 [1, 2, 4]
}

//26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。,

// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。,
func Task7() {
	method := func(intervals [][2]int) [][2]int {
		if len(intervals) == 0 {
			return [][2]int{}
		}
		j := 0
		for i := 1; i < len(intervals); i++ {
			if intervals[i][0] > intervals[j][1] {
				j++
				intervals[j] = intervals[i]
			}

			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		}
		return intervals[:j+1]

	}
	//{1,6},{3,5}
	fmt.Println(method([][2]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // 输出 [[1,6],[8,10],[15,18]]
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := 0
		numi := arr[i]
		for ; j < i; j++ {
			if numi < arr[j] {
				arr[i] = arr[j]
				break
			}
		}

		arr[j] = numi
	}

	fmt.Println(arr)
}

// 两数之和 ,
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func Task8() {
	method := func(nums []int, target int) []int {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
		return nil
	}
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(method(nums, target))
}
