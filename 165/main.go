package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1,".")
	v2 := strings.Split(version2,".")
	v1Len := len(v1)
	v2Len := len(v2)
	maxLen := v1Len
	if v2Len > v1Len {
		maxLen = v2Len
	}


	for i := 0; i < maxLen; i++ {
		value1 := 0
		if i >= v1Len {
			value1 = 0
		}else{
			value1,_ = strconv.Atoi(v1[i])
		}

		value2 := 0
		if i >= v2Len {
			value2 = 0
		}else{
			value2,_ = strconv.Atoi(v2[i])
		}

		if value1 > value2 {
			return 1
		}else if value1 < value2 {
			return -1
		}
	}




	return 0
}

func main() {
	fmt.Println(compareVersion("1.1.1","1.01.01"))
}
