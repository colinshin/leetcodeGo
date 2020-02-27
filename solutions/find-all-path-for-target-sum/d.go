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

func getPath(caps []int, relations []nodeRelation, s int) []string {
	fmt.Println(caps, relations, s)
	return nil
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	ints, err := readIntSlice(inputReader)
	if err != nil {
		return
	}
	m, s := ints[1], ints[2]
	capacity, err := readIntSlice(inputReader)
	if err != nil {
		return
	}

	nodeRelations := make([]nodeRelation, m)
	for i := 0; i < m; i++ {
		ints, err := readIntSlice(inputReader)
		if err != nil {
			return
		}
		nodeRelations[i].id = ints[0]
		nodeRelations[i].childs = ints[2:]
	}

	paths := getPath(capacity, nodeRelations, s)
	for i := range paths {
		fmt.Println(paths[i])
	}
}

type nodeRelation struct {
	id     int
	childs []int
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
