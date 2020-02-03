package powx_n

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func myPow(x float64, n int) float64 {
	// 先将n<0的情况转换为n为正数的情况
	if n < 0 {
		return pow(1/x, -n)
	}
	return pow(x, n)
}

/*
二分递归
*/
// n >= 0
func pow1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	halfPow := pow1(x, n/2)
	if n%2 == 1 {
		return halfPow * halfPow * x
	}
	return halfPow * halfPow
}

/*
二分迭代
来自于上边递归的思想

我们可以使用 n 的二进制表示来更好的理解该问题
我们定义result为结果，初始值为1
再用一个变量currentProduct记录x不断翻倍乘自身的结果
做一个这样的循环：
n不断折半（相当于二进制不断右移一位），每次折半currentProduct乘自身
如果n是奇数（相当于二进制最后一位是1），则result应该乘currentProduct;反之不用乘
*/
// n >= 0
func pow(x float64, n int) float64 {
	result := 1.0
	currentProduct := x
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			result *= currentProduct
		}
		currentProduct *= currentProduct
	}
	return result
}
