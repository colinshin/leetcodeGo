/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
删除数组里一个元素，是的数组有序；如果找不到这样的元素，输出-1
有序指递增或递减，但不要求严格递增或递减， 如1 2 2 3也算有序
如果s已经有序，直接输出-1
如果有多个可能的元素，输出最小的那个
*/
func delNum(s []int) int {
	n := len(s)
	if n < 2 {
		return -1
	}
	if n == 2 {
		return min(s[0], s[1])
	}
	peek, valley := 0, 0
	peekNum, valleyNum := 1, 1
	for i := 1; i < n-1; i++ {
		if s[i-1] < s[i] && s[i+1] < s[i] && s[i] > peek {
			peek = i
			peekNum++
		}
		if s[i-1] > s[i] && s[i+1] > s[i] && s[i] < valley {
			valley = i
			valleyNum++
		}
	}
	if peekNum > 1 || valleyNum > 1 {
		return -1
	}
	if peek == 0 && valley == n-1 || peek == n-1 && valley == 0 {
		return -1
	}
	if peek < valley && s[valley] >= s[peek-1] {

		return peek
	}
	if valley < peek && s[]
	return -1
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	_, err := readIntSlice(inputReader)
	if err != nil {
		return
	}

	s, err := readIntSlice(inputReader)
	if err != nil {
		return
	}
	fmt.Print(delNum(s))
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
