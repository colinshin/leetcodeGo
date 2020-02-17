package missing_number_in_arithmetic

/*
示例 1：

输入：arr = [5,7,11,13]
输出：9
解释：原来的数组是 [5,7,9,11,13]。
示例 2：

输入：arr = [15,13,12]
输出：14
解释：原来的数组是 [15,14,13,12]。


提示：

3 <= arr.length <= 1000
0 <= arr[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number-in-arithmetic-progression
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingNumber(arr []int) int {
	n := len(arr)
	d := (arr[n-1] - arr[0]) / n
	for i := 0; i < n-1; i++ {
		if arr[i]+d != arr[i+1] {
			return arr[i] + d
		}
	}
	return 0
}

func missingNumber1(arr []int) int {
	n := len(arr)
	sum := (arr[0] + arr[n-1]) * (n + 1) / 2
	for _, v := range arr {
		sum -= v
	}
	return sum
}
