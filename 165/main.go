package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNextValue(version string,begin,l int) (int,int)   {
	for begin  < l && version[begin] == '0'  {
		 begin++
	}
	if begin >= l {
		return 0,begin
	}

	res := 0

	for begin < l  {
		if version[begin] == '.' {
			for begin < l && version[begin] == '.' {
				begin++
			}
			break
		}else{
			res = res * 10 + int(version[begin] - '0')
			begin++
		}
	}

	return res,begin
}

func compareVersion(version1 string, version2 string) int {
	begin1 := 0
	begin2 := 0
	len1 := len(version1)
	len2 := len(version2)
	v1 := 0
	v2 := 0

	for begin1 < len1 || begin2 < len2 {
		v1,begin1 = getNextValue(version1,begin1,len1)
		v2,begin2 = getNextValue(version2,begin2,len2)

		if v1 > v2 {
			return 1
		}else if v2 > v1 {
			return -1
		}
	}

	return 0

}

func compareVersion2(version1 string, version2 string) int {
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

	fmt.Println(compareVersion2("1.1","1.1.0"))
}
