/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
给定一个矩阵，其中各个数字表示：
0 空位
1 石头
2 宝石
3 初始位置
4 结束位置
其中石头不可通过；每次可以向上下左右任意一个地方走一步
*/
/*
在步数限制内，最多能拿到几颗宝石，如果在规定限制内到不了结束位置，返回-1
*/
func getNumOfLightsOff(office [][]int, timeLimit int) int {

	return -1
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	ints, err := readIntSlice(inputReader)
	if err != nil {
		return
	}
	n := ints[0]

	ints, err = readIntSlice(inputReader)
	if err != nil {
		return
	}
	m := ints[0]
	office := make([][]int, n)
	for i := 0; i < n; i++ {
		office[i], err = readIntSlice(inputReader)
		if err != nil {
			return
		}
	}

	fmt.Print(getNumOfLightsOff(office, m))
}

func readIntSlice(reader *bufio.Reader) ([]int, error) {
	lineBuf, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf(err.Error())
	}

	lineBuf = strings.TrimRight(lineBuf, "\n")
	line := strings.Split(lineBuf, " ")
	var result []int
	for _, v := range line {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		result = append(result, int(i))
	}
	return result, nil
}
